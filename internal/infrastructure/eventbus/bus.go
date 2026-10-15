package eventbus

import (
	"log/slog"
	"reflect"
	"sync"
)

// Bus is an in-memory publish/subscribe bus. Handlers run asynchronously in their own goroutine,
// so events are lost if the process stops before a handler finishes.
type Bus struct {
	mu       sync.RWMutex
	handlers map[reflect.Type][]func(any)
}

func New() *Bus {
	return &Bus{
		handlers: make(map[reflect.Type][]func(any)),
	}
}

// Subscribe registers handler for events of type T. It is a function rather than a method
// because Go methods cannot have type parameters.
func Subscribe[T any](b *Bus, handler func(T)) {
	eventType := reflect.TypeFor[T]()

	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventType] = append(b.handlers[eventType], func(e any) {
		handler(e.(T))
	})
}

func (b *Bus) Publish(e any) {
	b.mu.RLock()
	handlers := b.handlers[reflect.TypeOf(e)]
	b.mu.RUnlock()

	for _, handler := range handlers {
		go run(handler, e)
	}
}

// A panic in a goroutine crashes the whole process, so a failing handler is recovered and logged instead.
func run(handler func(any), e any) {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("event handler panicked", "event", reflect.TypeOf(e).String(), "panic", r)
		}
	}()

	handler(e)
}
