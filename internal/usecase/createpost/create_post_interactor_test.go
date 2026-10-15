package createpost

import (
	"testing"

	"github.com/ShaimaaSabry/instagram/internal/domain/event"
)

type fakePublisher struct {
	published []any
}

func (f *fakePublisher) Publish(e any) {
	f.published = append(f.published, e)
}

func TestExecute_givenValidCommand_thenPublishesPostCreated(t *testing.T) {
	publisher := &fakePublisher{}
	interactor := NewInteractor(publisher)

	post, err := interactor.Execute(Command{
		Media:   []Media{{Url: "https://example.com/photo.jpg"}},
		Caption: "Hello",
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(publisher.published) != 1 {
		t.Fatalf("expected 1 event, got %d", len(publisher.published))
	}
	e, ok := publisher.published[0].(event.PostCreated)
	if !ok {
		t.Fatalf("expected event.PostCreated, got %T", publisher.published[0])
	}
	if e.PostID != post.Id() || e.Caption != "Hello" {
		t.Fatalf("unexpected event: %+v", e)
	}
}
