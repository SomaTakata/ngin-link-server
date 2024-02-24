package router

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func NewRouter(userController controller.UserController, linkController controller.LinkController) *gin.Engine {
	r := setupRouter()

	//health check
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	userRouter := r.Group("/users")
	{
		userRouter.GET("", userController.Get)
		userRouter.GET("/exists", userController.Exists)
		userRouter.POST("", userController.Create)
		//userRouter.PATCH("", userController.Update)
	}

	linkRouter := r.Group("/links")
	{
		linkRouter.GET("/:ngin-link-id", linkController.GetByNginLinkID)
		//linkRouter.PATCH("", linkController.Update)
		linkRouter.GET("/exchanged-ngin-links", linkController.GetExchangeHistory)
		linkRouter.POST("/exchanged-ngin-links/:ngin-link-id", linkController.CreateExchangeHistory)
	}

	return r
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("ALLOW_ORIGIN")},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		AllowCredentials: true,
	}))

	return r
}
