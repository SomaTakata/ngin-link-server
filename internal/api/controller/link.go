package controller

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/httpmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/usecase"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/clerkutil"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/httperror"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkController interface {
	GetByNginLinkID(ctx *gin.Context)
	Update(ctx *gin.Context)
	GetExchangeHistory(ctx *gin.Context)
	CreateExchangeHistory(ctx *gin.Context)
}

func NewLinkController(u usecase.LinkUsecase, client clerk.Client) LinkController {
	return &linkController{u, client}
}

type linkController struct {
	linkUsecase usecase.LinkUsecase
	client      clerk.Client
}

func (c linkController) GetByNginLinkID(ctx *gin.Context) {
	nginLinkID := ctx.Param("ngin-link-id")

	user, err := c.linkUsecase.GetByNginLinkID(nginLinkID)
	if err != nil && err.Error() == "record not found" {
		httperror.Handle(ctx, err, http.StatusNotFound)
		return
	}
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	nginLinkInfo := &httpmodel.NginLinkInfo{
		NginLink: &httpmodel.NginLink{
			NginLinkID:  user.NginLink.NginLinkID,
			SocialLinks: modelconverter.SocialLinksToHTTPModels(user.NginLink.SocialLinks),
		},
		Username:             user.Username,
		ProfileImageURL:      user.ProfileImageURL,
		Description:          user.Description,
		ProgrammingLanguages: user.ProgrammingLanguages,
		JobRole:              user.JobRole,
	}

	ctx.JSON(http.StatusOK, nginLinkInfo)
}

func (c linkController) Update(ctx *gin.Context) {
	clerkID, err := clerkutil.GetClerkID(ctx, c.client)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusUnauthorized)
		return
	}

	var socialLinksStruct httpmodel.SocialLinksStruct
	if err := ctx.BindJSON(&socialLinksStruct); err != nil {
		httperror.Handle(ctx, err, http.StatusBadRequest)
		return
	}
	reqSocialLinks := socialLinksStruct.SocialLinks

	socialLinks := modelconverter.SocialLinksFromHTTPModels(reqSocialLinks)
	newSocialLinks, err := c.linkUsecase.Update(clerkID, socialLinks)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, modelconverter.SocialLinksToStructHTTPModel(newSocialLinks))
}

func (c linkController) GetExchangeHistory(ctx *gin.Context) {
	clerkID, err := clerkutil.GetClerkID(ctx, c.client)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusUnauthorized)
		return
	}

	nginLinkExchangeHistory, err := c.linkUsecase.GetExchangeHistory(clerkID)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, modelconverter.NginLinkExchangeHistoryToHTTPModel(nginLinkExchangeHistory))
}

func (c linkController) CreateExchangeHistory(ctx *gin.Context) {
	clerkID, err := clerkutil.GetClerkID(ctx, c.client)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusUnauthorized)
		return
	}

	nginLinkID := ctx.Param("ngin-link-id")
	nginLinkExchangeHistory, err := c.linkUsecase.CreateExchangeHistory(clerkID, nginLinkID)
	if err != nil {
		httperror.Handle(ctx, err, http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, modelconverter.NginLinkExchangeHistoryToHTTPModel(nginLinkExchangeHistory))
}
