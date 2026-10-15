package fanoutpost

import "log/slog"

type Interactor struct {
	logger *slog.Logger
}

func NewInteractor(logger *slog.Logger) *Interactor {
	return &Interactor{
		logger: logger,
	}
}

type Command struct {
	PostID int
}

// Execute will add the post to its author's followers' feeds; users and followers do not exist yet, so it only logs.
func (i *Interactor) Execute(command Command) error {
	i.logger.Info(
		"fanning out post to followers' feeds",
		"postId", command.PostID,
	)

	return nil
}
