package eventbus

import (
	"testing"
	"time"
)

type eventA struct{ value string }
type eventB struct{}

func TestPublish_givenSubscriber_thenHandlerReceivesEvent(t *testing.T) {
	bus := New()
	received := make(chan eventA, 1)
	Subscribe(bus, func(e eventA) { received <- e })

	bus.Publish(eventA{value: "hello"})

	select {
	case e := <-received:
		if e.value != "hello" {
			t.Fatalf("expected hello, got %q", e.value)
		}
	case <-time.After(time.Second):
		t.Fatal("handler was not called")
	}
}

func TestPublish_givenSubscriberForOtherType_thenHandlerNotCalled(t *testing.T) {
	bus := New()
	called := make(chan struct{}, 1)
	Subscribe(bus, func(e eventB) { called <- struct{}{} })

	bus.Publish(eventA{})

	select {
	case <-called:
		t.Fatal("handler for eventB should not receive eventA")
	case <-time.After(50 * time.Millisecond):
	}
}

func TestPublish_givenPanickingHandler_thenOtherHandlersStillRun(t *testing.T) {
	bus := New()
	received := make(chan struct{}, 1)
	Subscribe(bus, func(e eventA) { panic("boom") })
	Subscribe(bus, func(e eventA) { received <- struct{}{} })

	bus.Publish(eventA{})

	select {
	case <-received:
	case <-time.After(time.Second):
		t.Fatal("second handler was not called")
	}
}
