package createpost

import (
	"time"

	"github.com/ShaimaaSabry/instagram/internal/domain/event"
	"github.com/ShaimaaSabry/instagram/internal/domain/model"
)

type eventPublisher interface {
	Publish(e any)
}

type Interactor struct {
	eventPublisher eventPublisher
}

func NewInteractor(eventPublisher eventPublisher) *Interactor {
	return &Interactor{
		eventPublisher: eventPublisher,
	}
}

type Media struct {
	Url string
}

type Command struct {
	Media   []Media
	Caption string
}

func (i *Interactor) Execute(command Command) (*model.Post, error) {
	media, err := createMedia(command)
	if err != nil {
		return nil, err
	}

	post, err := model.NewPost(
		media,
		command.Caption,
	)
	if err != nil {
		return nil, err
	}

	i.eventPublisher.Publish(event.PostCreated{
		PostID:     post.Id(),
		Caption:    post.Caption(),
		OccurredAt: time.Now(),
	})

	return &post, nil
}

func createMedia(command Command) ([]model.Media, error) {
	media := make([]model.Media, 0, len(command.Media))
	for _, m := range command.Media {
		mediaItem, err := model.NewMedia(m.Url)
		if err != nil {
			return nil, err
		}
		media = append(media, mediaItem)
	}

	return media, nil
}
