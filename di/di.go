package di

import (
	"github.com/daichi1002/go-graphql/adapters/repositories"
	"github.com/daichi1002/go-graphql/infra/db"
	"github.com/daichi1002/go-graphql/infra/envvars"
	"github.com/daichi1002/go-graphql/usecases"
	"github.com/daichi1002/go-graphql/usecases/user"
)

type Functions struct {
	GetUsers   usecases.GetUsersUsecase
	GetUser    usecases.GetUserUsecase
	CreateUser usecases.CreateUserUsecase
	UpdateUser usecases.UpdateUserUsecase
	DeleteUser usecases.DeleteUserUsecase
}

var functions *Functions

func Do(env envvars.EnvironmentVariablesInterface) {
	sqlHandler := db.NewSqlHandler(env)

	// repositories
	userRepository := repositories.NewUserRepository(sqlHandler)
	// usecases
	getUsersUsecase := user.NewGetUsersInteractor(userRepository)
	getUserUsecase := user.NewGetUserInteractor(userRepository)
	createUserUsecase := user.NewCreateUserInteractor(userRepository)
	updateUserUsecase := user.NewUpdateUserInteractor(userRepository)
	deleteUserUsecase := user.NewDeleteUserInteractor(userRepository)

	functions = &Functions{
		GetUsers:   getUsersUsecase,
		GetUser:    getUserUsecase,
		CreateUser: createUserUsecase,
		UpdateUser: updateUserUsecase,
		DeleteUser: deleteUserUsecase,
	}
}

func Provide() *Functions {
	return functions
}
