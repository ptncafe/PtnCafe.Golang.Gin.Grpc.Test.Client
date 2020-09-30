package app_error

import "github.com/pkg/errors"

type AppError struct {
	DetailError error
	Field string
	TypeError ErrorType
}

func NewAppErrorMessage(message string,typeError  ErrorType ) AppError{
	return AppError{DetailError: errors.New(message), TypeError: typeError}
}
func NewAppErrorFieldMessage(message string,typeError  ErrorType ,field string) *AppError{
	return &AppError{DetailError: errors.New(message), TypeError: typeError, Field:field}
}
func NewAppError(err error,typeError  ErrorType) AppError{
	return AppError{DetailError: err, TypeError: typeError}
}
func NewAppErrorNil(err error,typeError  ErrorType) *AppError{
	return &AppError{DetailError: err, TypeError: typeError}
}
func(appError *AppError) Error() string  {
	if appError == nil {
		return ""
	}
	if appError.DetailError != nil {
		return  appError.DetailError.Error()
	}
	return  "Đã có lỗi trong hệ thống"
}

func NewInvalidError(message string) AppError{
	return NewAppErrorMessage(message, BadRequestError)
}
func NewInvalidErrorNil(message string) *AppError{
	err := NewInvalidError(message)
	return &err
}
//func NewInvalidFieldError(message string,field string) AppError{
//	return NewAppErrorFieldMessage(message, BadRequestError, field)
//}
func NewInternalErrorMessage(message string) AppError{
	return NewAppErrorMessage(message, InternalError)
}
func NewInternalError(err error) AppError{
	return NewAppError(err, InternalError)
}
func NewInternalErrorNil(err error) *AppError{
	return NewAppErrorNil(err, InternalError)
}
type ErrorType int

const (
	InternalError ErrorType = 500
	BadRequestError ErrorType = 400
)