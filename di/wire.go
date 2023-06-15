//go:build wireinject
// +build wireinject

package di

import (
	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/usecases"
	"github.com/daichi1002/go-graphql/usecases/user"
	"github.com/google/wire"
)

func initGetUser() usecases.GetUserUsecase {
	wire.Build(
		repositories.NewUserRepository,
		user.NewGetUserInteractor,
	)
	return nil
}
