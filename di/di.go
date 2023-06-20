package di

import (
	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/infra/db"
	"github.com/daichi1002/go-graphql/infra/envvars"
	"github.com/daichi1002/go-graphql/usecases"
	"github.com/daichi1002/go-graphql/usecases/user"
)

type Functions struct {
	getUser usecases.GetUserUsecase
}

var functions *Functions

func Do(env envvars.EnvironmentVariablesInterface) {
	sqlHandler := db.NewSqlHandler(env)

	// repositories
	userRepository := repositories.NewUserRepository(sqlHandler)
	// usecases
	getUserUsecase := user.NewGetUserInteractor(userRepository)

	functions = &Functions{
		getUser: getUserUsecase,
	}
}

func Provide() *Functions {
	return functions
}
