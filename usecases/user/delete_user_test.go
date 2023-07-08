package user

import (
	"context"
	"fmt"
	"testing"

	mock_repositories "github.com/daichi1002/go-graphql/adapters/repositories/mock"
	"github.com/golang/mock/gomock"
)

func TestDeleteUserHandle(t *testing.T) {
	// パラメータ
	ctx := context.Background()
	userId := "1"

	// 期待するレスポンス
	deleteErr := fmt.Errorf("failed")

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
				ur.EXPECT().DeleteUser(ctx, userId).Return(nil)
			},
			expected: nil,
		},
		{
			name: "Failed",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().DeleteUser(ctx, userId).Return(deleteErr)
			},
			expected: deleteErr,
		},
	}

	// 各初期化処理
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repositories.NewMockUserRepository(ctrl)
	deleteUserInteractor := NewDeleteUserInteractor(mockUserRepository)
	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepository)

			err := deleteUserInteractor.Handle(ctx, userId)

			if err != tt.expected {
				t.Logf("err: %v", err)
				t.Logf("expected err: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
