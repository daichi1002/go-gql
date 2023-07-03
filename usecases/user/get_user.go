package user

import (
	"github.com/daichi1002/go-graphql/adapters/graph/repositories"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/usecases"
)

type GetUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUserInteractor(userRepository repositories.UserRepository) usecases.GetUserUsecase {
	return &GetUserInteractor{userRepository: userRepository}
}

func (interactor GetUserInteractor) Handle(userId string) (*model.User, error) {
	user, err := interactor.userRepository.GetUser(userId)

	if err != nil {
		return nil, err
	}

	return user, nil
}
