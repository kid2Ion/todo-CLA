package usecase

import (
	"errors"
	reflect "reflect"
	"testing"
	model "todo_CLA/domain/model"

	"github.com/golang/mock/gomock"
)

// go test -v

func TestView(t *testing.T) {
	tests := []struct {
		name     string
		expected []*model.Todo
		err      error
		wantErr  bool
	}{
		{
			name: "正常系",
			expected: []*model.Todo{
				{
					Id:        1,
					Task:      "test1",
					LimitDate: "あああ",
					Status:    true,
				},
				{
					Id:        9,
					Task:      "test2",
					LimitDate: "いいい",
					Status:    false,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:     "異常系",
			expected: nil,
			err:      errors.New("DB error"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockTodoRepository(ctrl)
			mock.EXPECT().FindAll().Return(tt.expected, tt.err)

			todoUsecase := NewTodoUsecase(mock)
			result, err := todoUsecase.View()

			if (err != nil) != tt.wantErr {
				t.Error("got err:", err)
			}
			for i, got := range result {
				want := tt.expected[i]
				t.Log(got)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got:\n%v\n\nwant:\n%v", result, want)
				}
			}
		})
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name     string
		expected []*model.Todo
		err      error
		wantErr  bool
	}{
		{
			name: "正常系",
			expected: []*model.Todo{
				{
					Id:        4,
					Task:      "歯磨き",
					LimitDate: "夜",
					Status:    false,
				},
			},
			err:     nil,
			wantErr: false,
		},
		{
			name:     "異常系",
			expected: nil,
			err:      errors.New("DB error"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := NewMockTodoRepository(ctrl)
			mock.EXPECT().Find("歯磨き").Return(tt.expected, tt.err)

			todoUsecase := NewTodoUsecase(mock)
			result, err := todoUsecase.Search("歯磨き")

			if (err != nil) != tt.wantErr {
				t.Error("got err:", err)
			}

			for i, got := range result {
				want := tt.expected[i]
				t.Log(got)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("got:\n%v\n\nwant:\n%v", result, want)
				}
			}
		})
	}
}
