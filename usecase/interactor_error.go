package usecase

import "fmt"

type TaskNameEmptyError struct {
	err error
}

func (e TaskNameEmptyError) Error() string {
	return fmt.Sprintf("task name must not be empty")
}

func (e TaskNameEmptyError) Unwrap() error {
	return e.err
}

type TaskNameAlreadyExistsError struct {
	name string
	err  error
}

func (e TaskNameAlreadyExistsError) Error() string {
	return fmt.Sprintf("task name `%s` is already exists", e.name)
}

func (e TaskNameAlreadyExistsError) Unwrap() error {
	return e.err
}
