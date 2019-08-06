package utils

type AppError struct {
	Code int
	Err  error
}

func (appError *AppError) Error() string {
	return appError.Err.Error()
}
