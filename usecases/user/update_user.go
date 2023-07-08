package user

import (
	"context"

	"github.com/daichi1002/go-graphql/adapters/repositories"
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
	err := interactor.userRepository.UpdateUser(ctx, input)

	if err != nil {
		return err
	}

	// TODO：user_idが存在しない場合はINVALID_PARAMETERのエラーを返す

	return nil
}
