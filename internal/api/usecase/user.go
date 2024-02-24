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
	user, err := u.userRepository.Get(clerkID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u userUsecase) Create(user *model.User) (*model.User, error) {
	newUser, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u userUsecase) Update(user *model.User) (*model.User, error) {
	//TODO
	return &model.User{}, nil
}
