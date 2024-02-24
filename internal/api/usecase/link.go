package usecase

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/repository"
)

type LinkUsecase interface {
	GetByNginLinkID(nginLinkID string) (*model.User, error)
	//Update
	//CreateExchangeHistory
}

func NewLinkUsecase(r repository.UserRepository) LinkUsecase {
	return &linkUsecase{r}
}

type linkUsecase struct {
	linkRepository repository.UserRepository
}

func (u linkUsecase) GetByNginLinkID(nginLinkID string) (*model.User, error) {
	return u.linkRepository.GetByNginLinkID(nginLinkID)
}
