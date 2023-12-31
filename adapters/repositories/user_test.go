package repositories

import (
	"context"
	"fmt"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/daichi1002/go-graphql/infra/db"
	mock_infra "github.com/daichi1002/go-graphql/infra/mock"
	"github.com/golang/mock/gomock"
)

func TestGetsUser(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	// 期待するレスポンス
	expected := []*model.User{
		{
			UserID:   "1",
			Name:     "test name1",
			Email:    "test1@xxx.go.jp",
			Password: "password1",
		},
		{
			UserID:   "2",
			Name:     "test name2",
			Email:    "test2@xxx.go.jp",
			Password: "password2",
		},
	}

	// テストケース
	tests := []struct {
		name     string
		expected []*model.User
		err      error
	}{
		{
			name:     "Success",
			expected: expected,
			err:      nil,
		},
		{
			name:     "Failed",
			expected: nil,
			err:      fmt.Errorf("failed"),
		},
	}

	// モック
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}

	defer database.Close()

	sqlHandler := db.NewSqlHandlerOfDB(database)
	repository := NewUserRepository(sqlHandler)

	// 期待するクエリ
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

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// DBモックのレスポンス設定
			rows := sqlmock.NewRows([]string{
				"user_id",
				"name",
				"email",
				"password",
			})
			for _, r := range tt.expected {
				rows.AddRow(
					r.UserID,
					r.Name,
					r.Email,
					r.Password,
				)
			}

			mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows).WillReturnError(tt.err)

			result, err := repository.GetUsers(ctx)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Logf("result: %v", result)
				t.Logf("expected result: %v", tt.expected)
				t.Error("Failed")
			}

			if err != tt.err {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
func TestGetUser(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	userId := "1"

	// 期待するレスポンス
	expected := &model.User{
		UserID:   userId,
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	// テストケース
	tests := []struct {
		name     string
		expected *model.User
		err      error
	}{
		{
			name:     "Success",
			expected: expected,
			err:      nil,
		},
		{
			name:     "Failed",
			expected: nil,
			err:      fmt.Errorf("failed"),
		},
	}

	// モック
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}

	defer database.Close()

	sqlHandler := db.NewSqlHandlerOfDB(database)
	repository := NewUserRepository(sqlHandler)

	// 期待するクエリ
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

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// DBモックのレスポンス設定
			rows := sqlmock.NewRows([]string{
				"user_id",
				"name",
				"email",
				"password",
			}).AddRow(
				expected.UserID,
				expected.Name,
				expected.Email,
				expected.Password,
			)

			mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(userId).WillReturnRows(rows).WillReturnError(tt.err)

			result, err := repository.GetUser(ctx, userId)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Logf("result: %v", result)
				t.Logf("expected result: %v", tt.expected)
				t.Error("Failed")
			}

			if err != tt.err {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	input := model.CreateUserInfo{
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	FakeNow("2023-01-01 00:00:00", "2023-12-31 00:00:00")
	currentTime := CustomNow()

	// 期待するレスポンス
	expected := fmt.Errorf("failed")

	// テストケース
	tests := []struct {
		name     string
		expected error
	}{
		{
			name:     "Success",
			expected: nil,
		},
		{
			name:     "Failed",
			expected: expected,
		},
	}

	// モック
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}

	defer database.Close()

	sqlHandler := db.NewSqlHandlerOfDB(database)
	repository := NewUserRepository(sqlHandler)
	ctrl := gomock.NewController(t)

	// 期待するクエリ
	query := `
		INSERT INTO users (
			name,
			email,
			password,
			created_at,
			updated_at
		) VALUES (?, ?, ?, ?, ?)
	`

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mock_infra.NewMockResult(ctrl)

			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					input.Name,
					input.Email,
					input.Password,
					currentTime,
					currentTime,
				).
				WillReturnResult(result).
				WillReturnError(tt.expected)

			err := repository.CreateUser(ctx, input)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	input := model.UpdateUserInfo{
		UserID:   "1",
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	FakeNow("2023-01-01 00:00:00", "2023-12-31 00:00:00")
	currentTime := CustomNow()

	// 期待するレスポンス
	expected := fmt.Errorf("failed")

	// テストケース
	tests := []struct {
		name     string
		expected error
	}{
		{
			name:     "Success",
			expected: nil,
		},
		{
			name:     "Failed",
			expected: expected,
		},
	}

	// モック
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}

	defer database.Close()

	sqlHandler := db.NewSqlHandlerOfDB(database)
	repository := NewUserRepository(sqlHandler)
	ctrl := gomock.NewController(t)

	// 期待するクエリ
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

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mock_infra.NewMockResult(ctrl)

			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					input.Name,
					input.Email,
					input.Password,
					currentTime,
					input.UserID,
				).
				WillReturnResult(result).
				WillReturnError(tt.expected)

			err := repository.UpdateUser(ctx, input)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	userId := "1"

	FakeNow("2023-01-01 00:00:00", "2023-12-31 00:00:00")
	currentTime := CustomNow()

	// 期待するレスポンス
	expected := fmt.Errorf("failed")

	// テストケース
	tests := []struct {
		name     string
		expected error
	}{
		{
			name:     "Success",
			expected: nil,
		},
		{
			name:     "Failed",
			expected: expected,
		},
	}

	// モック
	database, mock, err := sqlmock.New()
	if err != nil {
		t.Error("sqlmock not work")
	}

	defer database.Close()

	sqlHandler := db.NewSqlHandlerOfDB(database)
	repository := NewUserRepository(sqlHandler)
	ctrl := gomock.NewController(t)

	// 期待するクエリ
	query := `
	UPDATE
		users
	SET
		deleted_at = ?
	WHERE
		user_id = ?
	`

	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mock_infra.NewMockResult(ctrl)

			mock.ExpectExec(regexp.QuoteMeta(query)).
				WithArgs(
					currentTime,
					userId,
				).
				WillReturnResult(result).
				WillReturnError(tt.expected)

			err := repository.DeleteUser(ctx, userId)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
