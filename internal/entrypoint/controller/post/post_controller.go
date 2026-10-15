package post

import (
	"github.com/ShaimaaSabry/instagram/internal/domain/model"
	"github.com/ShaimaaSabry/instagram/internal/usecase/createpost"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createPostInteractor interface {
	Execute(command createpost.Command) (*model.Post, error)
}

type Controller struct {
	createPostInteractor createPostInteractor
}

func NewController(
	createPostInteractor createPostInteractor,
) *Controller {
	return &Controller{
		createPostInteractor,
	}
}

// CreatePost godoc
// @Summary Create a new photo post, carousel post, or video post.
// @Tags Posts
// @Param payload body createPostRequest true "Create post payload"
// @Success 201 {object} postDto
// @Router /v1/posts [post]
func (c *Controller) CreatePost(ctx *gin.Context) {
	request := createPostRequest{
		Media: []mediaDto{
			{Url: "hi"},
		},
		Caption: "Hello",
	}

	post, _ := c.createPostInteractor.Execute(request.toCommand())

	ctx.JSON(
		http.StatusCreated,
		newPostDto(post),
	)
}

func (c *Controller) GetProfile() {

}

func (c *Controller) searchPosts() {

}

func (c *Controller) getFeed() {

}
