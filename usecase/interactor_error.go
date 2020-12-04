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
