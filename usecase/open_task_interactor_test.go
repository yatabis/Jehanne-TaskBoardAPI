package usecase

import (
	"errors"
	"github.com/yatabis/Jehanne/TaskBoard/domain"
	"testing"

	"github.com/golang/mock/gomock"

	mock "github.com/yatabis/Jehanne/TaskBoard/mock_usecase"
)

func TestOpenTaskInteractor_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockOpenTaskRepository(ctrl)
	interactor := &OpenTaskInteractor{
		repository: mockRepository,
	}

	t.Run("TaskRepositoryがOpenTaskのリストを返すテスト", func(t *testing.T) {
		task1 := domain.OpenTask{ID: 1}
		task2 := domain.OpenTask{ID: 2}
		tasks := []*domain.OpenTask{&task1, &task2}
		mockRepository.EXPECT().FindAll().Return(tasks, nil)
		result, err := interactor.List()
		if err != nil {
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
		if len(result) != len(tasks) {
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
		for i := range result {
			if result[i] != tasks[i] {
				t.Errorf("result: %+v\nerr: %+v", result, err)
			}
		}
	})

	t.Run("TaskRepositoryがエラーを返すテスト", func(t *testing.T) {
		expectedErr := errors.New("TaskRepository からのエラー")
		mockRepository.EXPECT().FindAll().Return(nil, expectedErr)
		result, err := interactor.List()
		if (result != nil) || (err != expectedErr) {
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
	})

}
