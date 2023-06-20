package repositories

import (
	"github.com/daichi1002/go-graphql/adapters"
	"github.com/daichi1002/go-graphql/entities/model"
)

type UserRepositoryDependencies struct {
	sqlHandler adapters.SqlHandler
}

func NewUserRepository(sqlHandler adapters.SqlHandler) UserRepository {
	return &UserRepositoryDependencies{
		sqlHandler,
	}
}

func (dep *UserRepositoryDependencies) GetUser(userId string) (*model.User, error) {
	return nil, nil
}
