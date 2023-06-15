package usecases

import "github.com/daichi1002/go-graphql/entities/model"

type GetUserUsecase interface {
	Handle(userId string) (*model.User, error)
}
