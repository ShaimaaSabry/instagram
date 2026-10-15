package listener

import (
	"log/slog"

	"github.com/ShaimaaSabry/instagram/internal/domain/event"
	"github.com/ShaimaaSabry/instagram/internal/usecase/fanoutpost"
)

type fanOutPostInteractor interface {
	Execute(command fanoutpost.Command) error
}

type PostListener struct {
	fanOutPostInteractor fanOutPostInteractor
	logger               *slog.Logger
}

func NewPostListener(
	fanOutPostInteractor fanOutPostInteractor,
	logger *slog.Logger,
) *PostListener {
	return &PostListener{
		fanOutPostInteractor: fanOutPostInteractor,
		logger:               logger,
	}
}

func (l *PostListener) OnPostCreated(e event.PostCreated) {
	err := l.fanOutPostInteractor.Execute(fanoutpost.Command{
		PostID: e.PostID,
	})
	if err != nil {
		l.logger.Error(
			"failed to fan out post",
			"postId", e.PostID,
			"error", err,
		)
	}
}
