package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/entities"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/usecases"
)

type UpdateUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewUpdateUserInteractor(userRepository repositories.UserRepository) usecases.UpdateUserUsecase {
	return &UpdateUserInteractor{userRepository: userRepository}
}

func (interactor UpdateUserInteractor) Handle(ctx context.Context, input model.UpdateUserInfo) error {
	user, err := interactor.userRepository.GetUser(ctx, input.UserID)

	if err != nil {
		return err
	}

	if user == nil {
		return entities.INVALID_PARAMETER
	}

	err = interactor.userRepository.UpdateUser(ctx, input)

	if err != nil {
		return err
	}

	return nil
}
