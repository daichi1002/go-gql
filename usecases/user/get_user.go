package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/usecases"
)

type GetUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewGetUserInteractor(userRepository repositories.UserRepository) usecases.GetUserUsecase {
	return &GetUserInteractor{userRepository: userRepository}
}

func (interactor GetUserInteractor) Handle(ctx context.Context, userId string) (*model.User, error) {
	user, err := interactor.userRepository.GetUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	// TODO:userがnilの場合にnot_foundエラーを返す
	return user, nil
}
