package repositories

import "github.com/daichi1002/go-graphql/entities/model"

type UserRepository interface {
	GetUser(userId string) (*model.User, error)
	CreateUser(input model.CreateUserInfo) error
}
