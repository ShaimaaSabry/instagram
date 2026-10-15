package post

import "github.com/ShaimaaSabry/instagram/internal/usecase/createpost"

type createPostRequest struct {
	Media   []mediaDto
	Caption string `json:"caption"`
}

func (r createPostRequest) toCommand() createpost.Command {
	return createpost.Command{
		Media: func(in []mediaDto) []createpost.Media {
			out := make([]createpost.Media, len(in))
			for i, m := range in {
				out[i] = createpost.Media{Url: m.Url}
			}
			return out
		}(r.Media),
		Caption: r.Caption,
	}
}
