package di

import (
	"github.com/daichi1002/go-graphql/adapters/graph/repositories"
	"github.com/daichi1002/go-graphql/infra/db"
	"github.com/daichi1002/go-graphql/infra/envvars"
	"github.com/daichi1002/go-graphql/usecases"
	"github.com/daichi1002/go-graphql/usecases/user"
)

type Functions struct {
	GetUser    usecases.GetUserUsecase
	CreateUser usecases.CreateUserUsecase
}

var functions *Functions

func Do(env envvars.EnvironmentVariablesInterface) {
	sqlHandler := db.NewSqlHandler(env)

	// repositories
	userRepository := repositories.NewUserRepository(sqlHandler)
	// usecases
	getUserUsecase := user.NewGetUserInteractor(userRepository)
	createUserUsecase := user.NewCreateUserInteractor(userRepository)

	functions = &Functions{
		GetUser:    getUserUsecase,
		CreateUser: createUserUsecase,
	}
}

func Provide() *Functions {
	return functions
}
