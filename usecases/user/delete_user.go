package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/entities"
	"github.com/daichi1002/go-graphql/usecases"
)

type DeleteUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewDeleteUserInteractor(userRepository repositories.UserRepository) usecases.DeleteUserUsecase {
	return &DeleteUserInteractor{userRepository: userRepository}
}

func (interactor DeleteUserInteractor) Handle(ctx context.Context, userId string) error {
	user, err := interactor.userRepository.GetUser(ctx, userId)

	if err != nil {
		return err
	}

	if user == nil {
		return entities.INVALID_PARAMETER
	}

	err = interactor.userRepository.DeleteUser(ctx, userId)

	if err != nil {
		return err
	}

	return nil
}
