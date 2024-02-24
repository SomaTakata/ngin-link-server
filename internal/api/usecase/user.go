package usecase

import (
	"github.com/SomaTakata/ngin-link-server/internal/api/model"
	"github.com/SomaTakata/ngin-link-server/internal/api/repository"
)

type UserUsecase interface {
	Get(clerkID string) (*model.User, error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func (u userUsecase) Get(clerkID string) (*model.User, error) {
	return u.userRepository.Get(clerkID)
}

func (u userUsecase) Create(user *model.User) (*model.User, error) {
	return u.userRepository.Create(user)
}

func (u userUsecase) Update(user *model.User) (*model.User, error) {
	//TODO
	return &model.User{}, nil
}
