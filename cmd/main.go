package main

import (
	"log/slog"

	"github.com/ShaimaaSabry/instagram/internal/infrastructure/eventbus"
	"github.com/ShaimaaSabry/instagram/internal/interface/controller/post"
	"github.com/ShaimaaSabry/instagram/internal/interface/listener"
	"github.com/ShaimaaSabry/instagram/internal/usecase/createpost"
	"github.com/ShaimaaSabry/instagram/internal/usecase/fanoutpost"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// This import is REQUIRED so the generated docs package registers your spec.
	_ "github.com/ShaimaaSabry/instagram/docs"
)

// @title        Instagram APIs
// @version      1.0
func main() {
	postController := dependencyInjection()

	startServer(postController)
}

func dependencyInjection() *post.Controller {
	bus := eventbus.New()

	fanOutPostInteractor := fanoutpost.NewInteractor(slog.Default())
	postListener := listener.NewPostListener(fanOutPostInteractor, slog.Default())
	eventbus.Subscribe(bus, postListener.OnPostCreated)

	createPostInteractor := createpost.NewInteractor(bus)

	postController := post.NewController(createPostInteractor)

	return postController
}

func startServer(postController *post.Controller) {
	r := gin.Default()

	r.POST("/v1/posts", postController.CreatePost)

	setupSwagger(r)

	r.Run()
}

func setupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
