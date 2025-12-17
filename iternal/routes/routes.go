package routes

import (
	c "gobackend/iternal/controllers"
	"gobackend/iternal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("Before --> AA")
		ctx.Next()
		println("Alter --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("Before --> BB")
		ctx.Next()
		println("Alter --> BB")
	}
}

func CC(ctx *gin.Context) {
	println("Before --> CC")
	ctx.Next()
	println("Alter --> CC")
}

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthenMiddleware(), AA(), BB(), CC)
	v1 := r.Group("/v1/2025")
	{
		// Define a simple GET endpoint
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user", c.NewUserController().GetUserById)

	}
	return r
}
