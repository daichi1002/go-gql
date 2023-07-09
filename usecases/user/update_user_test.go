package user

import (
	"context"
	"fmt"
	"testing"

	mock_repositories "github.com/daichi1002/go-graphql/adapters/repositories/mock"
	"github.com/daichi1002/go-graphql/entities"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/golang/mock/gomock"
)

func TestUpdateUserHandle(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	input := model.UpdateUserInfo{
		UserID:   "1",
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	// 期待するレスポンス
	user := &model.User{
		UserID:   "1",
		Name:     "test name",
		Email:    "test@xxx.go.jp",
		Password: "password",
	}

	updateErr := fmt.Errorf("failed")

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
				ur.EXPECT().GetUser(ctx, input.UserID).Return(user, nil)
				ur.EXPECT().UpdateUser(ctx, input).Return(nil)
			},
			expected: nil,
		},
		{
			name: "Not Found",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUser(ctx, input.UserID).Return(nil, nil)
			},
			expected: entities.INVALID_PARAMETER,
		},
		{
			name: "Failed",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUser(ctx, input.UserID).Return(user, nil)
				ur.EXPECT().UpdateUser(ctx, input).Return(updateErr)
			},
			expected: updateErr,
		},
	}

	// 各初期化処理
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repositories.NewMockUserRepository(ctrl)
	updateUserInteractor := NewUpdateUserInteractor(mockUserRepository)
	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepository)

			err := updateUserInteractor.Handle(ctx, input)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
