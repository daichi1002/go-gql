//go:build wireinject
// +build wireinject

package di

import (
	"github.com/daichi1002/go-graphql/adapters"
	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/infra/db"
	"github.com/daichi1002/go-graphql/infra/envvars"
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

func initDB() adapters.SqlHandler {
	wire.Build(
		envvars.NewEnvvar,
		db.NewSqlHandler,
	)
	return nil
}
