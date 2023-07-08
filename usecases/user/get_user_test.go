package user

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	mock_repositories "github.com/daichi1002/go-graphql/adapters/repositories/mock"
	"github.com/daichi1002/go-graphql/entities/model"
	"github.com/golang/mock/gomock"
)

func TestGetUserHandle(t *testing.T) {
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
		name          string
		prepareMockFn func(
			ur *mock_repositories.MockUserRepository,
		)
		expected *model.User
	}{
		{
			name: "Success",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUser(ctx, userId).Return(expected, nil)
			},
			expected: expected,
		},
		{
			name: "Failed",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUser(ctx, userId).Return(nil, fmt.Errorf("failed"))
			},
			expected: nil,
		},
	}

	// 各初期化処理
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repositories.NewMockUserRepository(ctrl)
	getUserInteractor := NewGetUserInteractor(mockUserRepository)
	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepository)

			result, _ := getUserInteractor.Handle(ctx, userId)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Logf("result: %v", result)
				t.Logf("expected result: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
