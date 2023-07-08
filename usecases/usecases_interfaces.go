package usecases

import (
	"context"

	"github.com/daichi1002/go-graphql/entities/model"
)

type GetUsersUsecase interface {
	Handle(ctx context.Context) ([]*model.User, error)
}
type GetUserUsecase interface {
	Handle(ctx context.Context, userId string) (*model.User, error)
}

type CreateUserUsecase interface {
	Handle(ctx context.Context, input model.CreateUserInfo) error
}

type UpdateUserUsecase interface {
	Handle(ctx context.Context, input model.UpdateUserInfo) error
}

type DeleteUserUsecase interface {
	Handle(ctx context.Context, userId string) error
}
