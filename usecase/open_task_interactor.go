package usecase

import (
	"errors"
	"github.com/yatabis/Jehanne/TaskBoard/domain"
)

type OpenTaskInteractor struct {
	repository OpenTaskRepository
}

func (interactor *OpenTaskInteractor) List() ([]*domain.OpenTask, error) {
	tasks, err := interactor.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (interactor *OpenTaskInteractor) Add(options *domain.OpenTaskOptions) (*domain.OpenTask, error) {
	if options.Name == "" {
		return nil, TaskNameEmptyError{}
	}
	now := domain.Now()
	today := domain.Today()
	options.Category.Init()
	if options.Repeating {
		options.WorkOn.Init(today)
		options.Deadline.Init(today)
	} else {
		options.WorkOn.Init(domain.Date{})
		options.Deadline.Init(domain.Date{})
	}
	options.Status.Init()
	task := domain.OpenTask{
		ID:              0,
		Name:            options.Name,
		Category:        options.Category,
		Repeating:       options.Repeating,
		WorkOn:          options.WorkOn,
		Deadline:        options.Deadline,
		Status:          options.Status,
		PerformanceTime: 0,
		CreatedAt:       now,
		UpdatedAt:       now,
	}
	result, err := interactor.repository.Save(&task)
	if err != nil {
		return nil, err
	}
	return result, nil
}
