package post

import "github.com/ShaimaaSabry/instagram/internal/domain/model"

type mediaDto struct {
	Url string `json:"url"`
}

func newMediaDto(media model.Media) mediaDto {
	return mediaDto{
		Url: media.Url(),
	}
}

type postDto struct {
	Id      int
	Media   []mediaDto
	Caption string
}

func newPostDto(post *model.Post) postDto {
	return postDto{
		Id: post.Id(),
		Media: func(in []model.Media) []mediaDto {
			out := make([]mediaDto, len(in))
			for i, m := range in {
				out[i] = newMediaDto(m)
			}
			return out
		}(post.Media()),
		Caption: post.Caption(),
	}
}
