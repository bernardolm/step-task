package errors

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type AppError struct {
	message string
	fields  log.Fields
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Fields() log.Fields {
	return e.fields
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) WithField(k string, v any) {
	e.fields[k] = v
}

func New(message string, args ...any) *AppError {
	return &AppError{
		fields:  log.Fields{},
		message: fmt.Sprintf(message, args...),
	}
}
