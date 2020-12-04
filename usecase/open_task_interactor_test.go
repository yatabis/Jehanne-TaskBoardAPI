package usecase

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	mock "github.com/yatabis/Jehanne/TaskBoard/mock_usecase"

	"github.com/yatabis/Jehanne/TaskBoard/domain"
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
			return
		}
		if len(result) != len(tasks) {
			t.Errorf("result: %+v\nerr: %+v", result, err)
			return
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

func TestOpenTaskInteractor_Add(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepository := mock.NewMockOpenTaskRepository(ctrl)
	interactor := &OpenTaskInteractor{
		repository: mockRepository,
	}
	domain.FixedTime = time.Now()
	now := domain.Now()
	today := domain.Today()

	t.Run("すべてのオプションを設定するテスト", func(t *testing.T) {
		name := "タスク名"
		category := domain.TaskCategory("カテゴリー")
		repeating := true
		workOn, err := domain.NewDate("2020-12-28")
		if err != nil {
			t.Errorf("%+v", err)
			return
		}
		deadline, err := domain.NewDate("2020-12-31")
		if err != nil {
			t.Errorf("%+v", err)
			return
		}
		status := domain.TaskStatusUnowned

		options := domain.OpenTaskOptions{
			Name:      name,
			Category:  category,
			Repeating: repeating,
			WorkOn:    workOn,
			Deadline:  deadline,
			Status:    status,
		}
		set := domain.OpenTask{
			ID:              0,
			Name:            name,
			Category:        category,
			Repeating:       repeating,
			WorkOn:          workOn,
			Deadline:        deadline,
			Status:          status,
			PerformanceTime: 0,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		saved := set
		saved.ID = 1
		mockRepository.EXPECT().FindByName(name).Return(nil, TaskNameNotFoundError{name: name})
		mockRepository.EXPECT().Save(&set).Return(&saved, nil)
		result, err := interactor.Add(&options)
		if err != nil {
			t.Errorf("err%+v", err)
			return
		}
		if result != &saved {
			g := reflect.ValueOf(*result)
			e := reflect.ValueOf(saved)
			err := ""
			for i := 0; i < g.Type().NumField(); i++ {
				f := g.Type().Field(i).Name
				gi := g.Field(i).Interface()
				ei := e.Field(i).Interface()
				if gi != ei {
					err += fmt.Sprintf("field `%s`\nGot: %v\nWant: %v\n", f, gi, ei)
				}
			}
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
	})

	t.Run("可能な限りオプションを省略するテスト", func(t *testing.T) {
		name := "タスク名"
		options := domain.OpenTaskOptions{
			Name: name,
		}
		set := domain.OpenTask{
			ID:              0,
			Name:            name,
			Category:        "Inbox",
			Repeating:       false,
			WorkOn:          domain.Date{},
			Deadline:        domain.Date{},
			Status:          domain.TaskStatusWaiting,
			PerformanceTime: 0,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		saved := set
		saved.ID = 2
		mockRepository.EXPECT().FindByName(name).Return(nil, TaskNameNotFoundError{name: name})
		mockRepository.EXPECT().Save(&set).Return(&saved, nil)
		result, err := interactor.Add(&options)
		if err != nil {
			t.Errorf("err%+v", err)
			return
		}
		if result != &saved {
			g := reflect.ValueOf(*result)
			e := reflect.ValueOf(saved)
			err := ""
			for i := 0; i < g.Type().NumField(); i++ {
				f := g.Type().Field(i).Name
				gi := g.Field(i).Interface()
				ei := e.Field(i).Interface()
				if gi != ei {
					err += fmt.Sprintf("field `%s`\nGot: %v\nWant: %v\n", f, gi, ei)
				}
			}
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
	})

	t.Run("可能な限りオプションを省略した定常タスクのテスト", func(t *testing.T) {
		name := "タスク名"
		repeating := true
		options := domain.OpenTaskOptions{
			Name:      name,
			Repeating: repeating,
		}
		set := domain.OpenTask{
			ID:              0,
			Name:            name,
			Category:        "Inbox",
			Repeating:       repeating,
			WorkOn:          today,
			Deadline:        today,
			Status:          domain.TaskStatusWaiting,
			PerformanceTime: 0,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		saved := set
		saved.ID = 3
		mockRepository.EXPECT().FindByName(name).Return(nil, TaskNameNotFoundError{name: name})
		mockRepository.EXPECT().Save(&set).Return(&saved, nil)
		result, err := interactor.Add(&options)
		if err != nil {
			t.Errorf("err%+v", err)
			return
		}
		if result != &saved {
			g := reflect.ValueOf(*result)
			e := reflect.ValueOf(saved)
			err := ""
			for i := 0; i < g.Type().NumField(); i++ {
				f := g.Type().Field(i).Name
				gi := g.Field(i).Interface()
				ei := e.Field(i).Interface()
				if gi != ei {
					err += fmt.Sprintf("field `%s`\nGot: %v\nWant: %v\n", f, gi, ei)
				}
			}
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
	})

	t.Run("タスク名を省略するテスト", func(t *testing.T) {
		options := domain.OpenTaskOptions{}
		result, err := interactor.Add(&options)
		if result != nil {
			t.Errorf("result is expexted to be nil but got %+v\n", result)
			return
		}
		expectedErr := TaskNameEmptyError{}
		if !errors.As(err, &expectedErr) {
			t.Errorf("err is expected to be %T but got %T\n", err, expectedErr)
		}
	})

	t.Run("タスク名が重複しているテスト", func(t *testing.T) {
		name := "タスク名"
		options := domain.OpenTaskOptions{
			Name: name,
		}
		task := domain.OpenTask{
			ID:              1,
			Name:            name,
			Category:        domain.TaskCategoryInbox,
			Repeating:       false,
			WorkOn:          today,
			Deadline:        today,
			Status:          domain.TaskStatusWaiting,
			PerformanceTime: 0,
			CreatedAt:       now,
			UpdatedAt:       now,
		}
		mockRepository.EXPECT().FindByName(name).Return(&task, nil)
		result, err := interactor.Add(&options)
		if result != nil {
			t.Errorf("result is expexted to be nil but got %+v\n", result)
			return
		}
		expectedErr := TaskNameAlreadyExistsError{name: name}
		if !errors.As(err, &expectedErr) {
			t.Errorf("err is expected to be %T but got %T\n", err, expectedErr)
		}
	})

	t.Run("TaskRepositoryがエラーを返すテスト", func(t *testing.T) {
		name := "タスク名"
		options := domain.OpenTaskOptions{
			Name: name,
		}
		expectedErr := errors.New("TaskRepository からのエラー")
		mockRepository.EXPECT().FindByName(name).Return(nil, expectedErr)
		result, err := interactor.Add(&options)
		if (result != nil) || (err != expectedErr) {
			t.Errorf("result: %+v\nerr: %+v", result, err)
		}
	})
	domain.FixedTime = time.Time{}
}
