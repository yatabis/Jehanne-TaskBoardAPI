//go:generate mockgen -source=open_task_repository.go -destination ../mock_usecase/mock_open_task_repository.go

package usecase

import "github.com/yatabis/Jehanne/TaskBoard/domain"

type OpenTaskRepository interface {
	FindAll() ([]*domain.OpenTask, error)
	Save(task *domain.OpenTask) (*domain.OpenTask, error)
}
