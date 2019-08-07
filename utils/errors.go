package utils

import "github.com/pkg/errors"

type AppError struct {
	Code int
	Err  error
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}

func GetAppError(err error, errMsg string, code int) *AppError {
	return &AppError{
		Code: code,
		Err:  errors.Wrap(err, errMsg),
	}

}
