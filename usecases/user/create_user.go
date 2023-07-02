package user

import (
	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/usecases"
)

type CreateUserInteractor struct {
	userRepository repositories.UserRepository
}

func NewCreateUserInteractor(userRepository repositories.UserRepository) usecases.CreateUserUsecase {
	return &CreateUserInteractor{userRepository: userRepository}
}

func (interactor CreateUserInteractor) Handle(input model.CreateUserInfo) error {
	err := interactor.userRepository.CreateUser(input)

	if err != nil {
		return err
	}

	return nil
}
