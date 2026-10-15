package listener

import (
	"errors"
	"log/slog"
	"testing"

	"github.com/ShaimaaSabry/instagram/internal/domain/event"
	"github.com/ShaimaaSabry/instagram/internal/usecase/fanoutpost"
)

type fakeFanOutPostInteractor struct {
	commands []fanoutpost.Command
	err      error
}

func (f *fakeFanOutPostInteractor) Execute(command fanoutpost.Command) error {
	f.commands = append(f.commands, command)
	return f.err
}

func TestOnPostCreated_thenExecutesFanOutWithPostId(t *testing.T) {
	interactor := &fakeFanOutPostInteractor{}
	listener := NewPostListener(interactor, slog.Default())

	listener.OnPostCreated(event.PostCreated{PostID: 42})

	if len(interactor.commands) != 1 || interactor.commands[0].PostID != 42 {
		t.Fatalf("expected one command for post 42, got %+v", interactor.commands)
	}
}

func TestOnPostCreated_givenFanOutFails_thenDoesNotPanic(t *testing.T) {
	interactor := &fakeFanOutPostInteractor{err: errors.New("boom")}
	listener := NewPostListener(interactor, slog.Default())

	listener.OnPostCreated(event.PostCreated{PostID: 42})
}
