package errors

import (
	"fmt"
)

type Action string

const (
	msgInternalError = "Oops! Something went wrong while"
)

type BusinessLogicError interface {
	Error() string
	toAPIError() ApiError
}

type InternalError struct {
	error
}

type NotFoundError struct {
	error
}

type BadRequestError struct {
	error
}

func NewNotFoundError(message string) BusinessLogicError {
	return &NotFoundError{error: fmt.Errorf(message)}
}

func NewInternalError(message string) BusinessLogicError {
	return &InternalError{error: fmt.Errorf(message)}
}

func NewBadRequestError(message string) BusinessLogicError {
	return &InternalError{error: fmt.Errorf(message)}
}

func (e *InternalError) toAPIError() ApiError {
	return NewInternalServerApiError(e.Error(), nil)
}

func (e *NotFoundError) toAPIError() ApiError {
	return NewNotFoundApiError(e.Error())
}

func (e *BadRequestError) toAPIError() ApiError {
	return NewBadRequestApiError(e.Error())
}

func GetAPIError(err error) ApiError {
	if businessError, ok := err.(BusinessLogicError); ok {
		return businessError.toAPIError()
	}

	return NewInternalServerApiError(msgInternalError, nil)
}

