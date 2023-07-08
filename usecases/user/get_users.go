package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/usecases"
)

type GetUsersInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUsersInteractor(userRepository repositories.UserRepository) usecases.GetUsersUsecase {
	return &GetUsersInteractor{userRepository: userRepository}
}

func (interactor GetUsersInteractor) Handle(ctx context.Context) ([]*model.User, error) {
	user, err := interactor.userRepository.GetUsers(ctx)

	if err != nil {
		return nil, err
	}

	return user, nil
}
