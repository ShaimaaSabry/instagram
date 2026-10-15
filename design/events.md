# Events

## Current design

Use cases publish domain events, and listeners react to them in the background.

```
POST /v1/posts → post.Controller → createpost.Interactor ──Publish(PostCreated)──► eventbus.Bus
                                                                                       │ goroutine
                                                                                       ▼
                                          fanoutpost.Interactor ◄── listener.PostListener.OnPostCreated
```

| Piece | Location | Role |
|---|---|---|
| `event.PostCreated` | `internal/domain/event` | Plain data describing what happened |
| `eventPublisher` interface | `internal/usecase/createpost` | What the use case needs; it does not know about the bus |
| `eventbus.Bus` | `internal/infrastructure/eventbus` | In-memory pub/sub; each handler runs in its own goroutine, panics are recovered and logged |
| `listener.PostListener` | `internal/entrypoint/listener` | Inbound adapter, like a controller: converts the event into a command and calls a use case |
| `fanoutpost.Interactor` | `internal/usecase/fanoutpost` | The reaction's logic (only logs until users and followers exist) |

Listeners are registered explicitly in `cmd/main.go` with `eventbus.Subscribe`.

### Limitations

- Events live only in memory: a crash or restart while a handler runs loses the event.
- Nothing retries a failed handler; the error is only logged.
- The server has no graceful shutdown, so in-flight handlers are cut off when the process stops.

## Later, if needed: Watermill

Go's standard library has no event bus, which is why `eventbus.Bus` is hand-written. If the limitations above start to matter, replace it with [Watermill](https://github.com/ThreeDotsLabs/watermill) (ThreeDotsLabs):

- It is the most established Go library for this.
- Its `gochannel` pub/sub works in memory, and delivery can later be switched to Kafka, RabbitMQ or SQL-backed delivery without changing handlers.
- Its router has retry and recovery middleware.
- It is the closest Go gets to Spring events plus Spring Modulith (event outbox and republishing).

The switch should only touch `internal/infrastructure` and the wiring in `cmd/main.go`: use cases depend on the `eventPublisher` interface, and listeners are plain functions.
