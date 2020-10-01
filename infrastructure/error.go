package infrastructure

import "fmt"

type AppErr struct {
	Level   ErrLevel
	Code    int
	Message string
	err     error
}

type ErrLevel string

func (err *AppErr) Error() string {
	return fmt.Sprintf("%s", err.Message)
}
