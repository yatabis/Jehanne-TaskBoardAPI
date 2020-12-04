package usecase

import "fmt"

type RepositoryError interface {
	Error() string
	Unwrap() error
}

type TaskNameNotFoundError struct {
	name string
	err  error
}

func (e TaskNameNotFoundError) Error() string {
	return fmt.Sprintf("task name `%s` is not found", e.name)
}

func (e TaskNameNotFoundError) Unwrap() error {
	return e.err
}
