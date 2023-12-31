package repositories

import (
	"context"

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

func (dep *UserRepositoryDependencies) GetUsers(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	query := `
	SELECT
		user_id,
		name,
		email,
		password
	FROM
		users
	WHERE
		deleted_at IS NULL
	`

	row, err := dep.sqlHandler.Query(ctx, query)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	for row.Next() {
		user := &model.User{}

		row.Scan(
			&user.UserID,
			&user.Name,
			&user.Email,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (dep *UserRepositoryDependencies) GetUser(ctx context.Context, userId string) (*model.User, error) {
	var user model.User

	query := `
	SELECT
		user_id,
		name,
		email,
		password
	FROM
		users
	WHERE
		user_id = ?
	AND deleted_at IS NULL
	`

	row, err := dep.sqlHandler.Query(ctx, query, userId)

	if err != nil {
		return nil, err
	}

	defer row.Close()

	// 取得0件の場合
	if !row.Next() {
		return nil, nil
	}

	err = row.Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (dep *UserRepositoryDependencies) CreateUser(ctx context.Context, input model.CreateUserInfo) error {
	currentTime := CustomNow()
	// create処理
	query := `
		INSERT INTO users (
			name,
			email,
			password,
			created_at,
			updated_at
		) VALUES (?, ?, ?, ?, ?)
	`

	_, err := dep.sqlHandler.Execute(
		ctx,
		query,
		input.Name,
		input.Email,
		input.Password,
		currentTime,
		currentTime,
	)

	if err != nil {
		return err
	}

	return nil
}

func (dep *UserRepositoryDependencies) UpdateUser(ctx context.Context, input model.UpdateUserInfo) error {
	currentTime := CustomNow()
	// update処理
	query := `
	UPDATE
		users
	SET
		name = ?,
		email = ?,
		password = ?,
		updated_at = ?
	WHERE
		user_id = ?
	`

	_, err := dep.sqlHandler.Execute(
		ctx,
		query,
		input.Name,
		input.Email,
		input.Password,
		currentTime,
		input.UserID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (dep *UserRepositoryDependencies) DeleteUser(ctx context.Context, userId string) error {
	currentTime := CustomNow()
	// update処理
	query := `
	UPDATE
		users
	SET
		deleted_at = ?
	WHERE
		user_id = ?
	`

	_, err := dep.sqlHandler.Execute(
		ctx,
		query,
		currentTime,
		userId,
	)

	if err != nil {
		return err
	}

	return nil
}
