package domain

import "fmt"

type ErrNotFound struct {
	EntityType string
	EntityID   string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s with id %s not found", e.EntityType, e.EntityID)
}

type ErrInvalidEntity struct {
	EntityType string
	Message    string
}

func (e *ErrInvalidEntity) Error() string {
	return fmt.Sprintf("Invalid %s: %s", e.EntityType, e.Message)
}
