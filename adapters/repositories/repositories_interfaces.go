package repositories

import (
	"context"

	"github.com/daichi1002/go-graphql/entities/model"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	GetUser(ctx context.Context, userId string) (*model.User, error)
	CreateUser(ctx context.Context, input model.CreateUserInfo) error
}

type TxRepository interface {
	DoInTx(f func(ctx context.Context) error) error
}
