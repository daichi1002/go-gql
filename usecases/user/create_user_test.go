package user

import (
	"context"
	"fmt"
	"testing"

	mock_repositories "github.com/daichi1002/go-graphql/adapters/repositories/mock"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/golang/mock/gomock"
)

func TestCreateUserHandle(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	input := model.CreateUserInfo{
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	// 期待するレスポンス
	createErr := fmt.Errorf("failed")

	// テストケース
	tests := []struct {
		name          string
		prepareMockFn func(
			ur *mock_repositories.MockUserRepository,
		)
		expected error
	}{
		{
			name: "Success",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().CreateUser(ctx, input).Return(nil)
			},
			expected: nil,
		},
		{
			name: "Failed",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().CreateUser(ctx, input).Return(createErr)
			},
			expected: createErr,
		},
	}

	// 各初期化処理
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repositories.NewMockUserRepository(ctrl)
	createuserInteractor := NewCreateUserInteractor(mockUserRepository)
	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepository)

			err := createuserInteractor.Handle(ctx, input)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
