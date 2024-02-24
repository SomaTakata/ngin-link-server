package controller

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/reqmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/usecase"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/clerkutil"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/httperror"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

func NewUserController(u usecase.UserUsecase, client clerk.Client) UserController {
	return &userController{u, client}
}

type userController struct {
	userUsecase usecase.UserUsecase
	client      clerk.Client
}

func (c userController) Get(ctx *gin.Context) {
	clerkID, err := clerkutil.GetClerkID(ctx, c.client)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusUnauthorized)
		return
	}

	user, err := c.userUsecase.Get(clerkID)
	if err != nil && err.Error() == "record not found" {
		httperror.Handle(ctx, err, http.StatusNotFound)
		return
	}
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c userController) Create(ctx *gin.Context) {
	var createUser reqmodel.CreateUser
	if err := ctx.ShouldBindJSON(&createUser); err != nil {
		httperror.Handle(ctx, err, http.StatusBadRequest)
		return
	}

	user := modelconverter.UserFromCreateUserReqModel(&createUser)

	clerkID, err := clerkutil.GetClerkID(ctx, c.client)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusUnauthorized)
		return
	}

	user.ClerkID = clerkID

	newUser, err := c.userUsecase.Create(user)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, newUser)
}

func (c userController) Update(ctx *gin.Context) {

}
