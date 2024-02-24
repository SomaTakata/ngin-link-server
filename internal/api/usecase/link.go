package usecase

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/repository"
)

type LinkUsecase interface {
	GetByNginLinkID(nginLinkID string) (*model.User, error)
	//Update
	GetExchangeHistory(clerkID string) (*model.NginLinkExchangeHistory, error)
	CreateExchangeHistory(clerkID string, nginLinkID string) (*model.NginLinkExchangeHistory, error)
}

func NewLinkUsecase(lr repository.LinkRepository, ur repository.UserRepository) LinkUsecase {
	return &linkUsecase{lr, ur}
}

type linkUsecase struct {
	linkRepository repository.LinkRepository
	userRepository repository.UserRepository
}

func (u linkUsecase) GetByNginLinkID(nginLinkID string) (*model.User, error) {
	//ほぼuser情報そのものを使うので、userRepositoryから取得する
	return u.userRepository.GetByNginLinkID(nginLinkID)
}

func (u linkUsecase) GetExchangeHistory(clerkID string) (*model.NginLinkExchangeHistory, error) {
	return u.linkRepository.GetExchangeHistory(clerkID)
}

func (u linkUsecase) CreateExchangeHistory(clerkID string, nginLinkID string) (*model.NginLinkExchangeHistory, error) {
	return u.linkRepository.CreateExchangeHistory(clerkID, nginLinkID)
}
