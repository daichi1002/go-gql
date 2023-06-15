package repositories

import "github.com/daichi1002/go-graphql/entities/model"

type UserRepositoryDependencies struct {
	// sqlHandle

}

func NewUserRepository() UserRepository {
	return &UserRepositoryDependencies{}
}

func (dep *UserRepositoryDependencies) GetUser(userId string) (*model.User, error) {
	return nil, nil
}
