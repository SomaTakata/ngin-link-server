package controller

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/resmodel"
	"github.com/SomaTakata/ngin-link-server/internal/api/usecase"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/httperror"
	"github.com/SomaTakata/ngin-link-server/internal/api/util/modelconverter"
	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LinkController interface {
	GetByNginLinkID(ctx *gin.Context)
	Update(ctx *gin.Context)
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

	nginLinkInfo := resmodel.NginLinkInfo{
		NginLink: &resmodel.NginLink{
			NginLinkID:  user.NginLink.NginLinkID,
			SocialLinks: modelconverter.SocialLinksToResModel(user.NginLink.SocialLinks),
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

}

func (c linkController) CreateExchangeHistory(ctx *gin.Context) {

}
