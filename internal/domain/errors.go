package domain

import "fmt"

type (
	ValidationError struct {
		Field   string
		Message string
	}

	NotFoundError struct {
		Entity string
		ID     any
	}
)

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %v not found", e.Entity, e.ID)
}
