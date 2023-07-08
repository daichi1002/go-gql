package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/usecases"
)

type DeleteUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewDeleteUserInteractor(userRepository repositories.UserRepository) usecases.DeleteUserUsecase {
	return &DeleteUserInteractor{userRepository: userRepository}
}

func (interactor DeleteUserInteractor) Handle(ctx context.Context, userId string) error {
	err := interactor.userRepository.DeleteUser(ctx, userId)

	if err != nil {
		return err
	}

	// TODO：user_idが存在しない場合はINVALID_PARAMETERのエラーを返す

	return nil
}
