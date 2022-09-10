package mock_repository

import (
	reflect "reflect"
	"testing"
	model "todo_CLA/domain/model"
	"todo_CLA/usecase"

	"github.com/golang/mock/gomock"
)

func TestView(t *testing.T) {
	// mockの呼び出しを管理するcontroller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected []*model.Todo
	var err error

	// mockの生成
	mock := NewMockTodoRepository(ctrl)
	mock.EXPECT().FindAll().Return(expected, err)

	todoUsecase := usecase.NewTodoUsecase(mock)
	result, err := todoUsecase.View()

	if err != nil {
		t.Error("Actual FindAll() is not same as expected")
	}

	if !reflect.DeepEqual(result, expected) {
		t.Error("Actual FindAll() is not same as expected")
	}

}
