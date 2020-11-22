package usecase

import (
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
