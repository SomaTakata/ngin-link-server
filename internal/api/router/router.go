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
		userRouter.POST("", userController.Create)
		//userRouter.PATCH("", userController.Update)
	}

	linkRouter := r.Group("/links")
	{
		linkRouter.GET("/:ngin-link-id", linkController.GetByNginLinkID)
		//linkRouter.PATCH("", linkController.Update)
		//linkRouter.POST("/exchange-ngin-link/:ngin-link-id", linkController.ExchangeHistoryCreate)
	}

	return r
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ALLOW_ORIGIN")},
		AllowCredentials: true,
	}))

	return r
}
