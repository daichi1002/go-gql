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

func TestGetUsersHandle(t *testing.T) {
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
		name          string
		prepareMockFn func(
			ur *mock_repositories.MockUserRepository,
		)
		expected []*model.User
	}{
		{
			name: "Success",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUsers(ctx).Return(expected, nil)
			},
			expected: expected,
		},
		{
			name: "Failed",
			prepareMockFn: func(
				ur *mock_repositories.MockUserRepository,
			) {
				ur.EXPECT().GetUsers(ctx).Return(nil, fmt.Errorf("failed"))
			},
			expected: nil,
		},
	}

	// 各初期化処理
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepository := mock_repositories.NewMockUserRepository(ctrl)
	getUsersInteractor := NewGetUsersInteractor(mockUserRepository)
	// テスト実行
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareMockFn(mockUserRepository)

			result, _ := getUsersInteractor.Handle(ctx)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Logf("result: %v", result)
				t.Logf("expected result: %v", tt.expected)
				t.Error("Failed")
			}
		})
	}
}
