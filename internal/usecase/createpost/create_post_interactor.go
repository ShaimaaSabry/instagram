package createpost

import "github.com/ShaimaaSabry/instagram/internal/domain/model"

type Interactor struct {
}

func NewInteractor() *Interactor {
	return &Interactor{}
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

	return &post, err
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
