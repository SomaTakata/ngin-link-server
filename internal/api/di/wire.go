//go:build wireinject
// +build wireinject

package di

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/controller"
	"github.com/SomaTakata/ngin-link-server/internal/api/db"
	"github.com/SomaTakata/ngin-link-server/internal/api/repository"
	"github.com/SomaTakata/ngin-link-server/internal/api/router"
	"github.com/SomaTakata/ngin-link-server/internal/api/usecase"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/clerkutil"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func Wire() *gin.Engine {
	wire.Build(
		db.NewDB,
		clerkutil.NewClerkClient,
		repository.NewUserRepository,
		usecase.NewUserUsecase,
		controller.NewUserController,
		router.NewRouter,
	)
	return &gin.Engine{}
}
