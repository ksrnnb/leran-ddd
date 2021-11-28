package errs

import "net/http"

type AppError struct {
	code    int
	message string
}

func NewAppError(code int, message string) *AppError {
	if code < 100 || code >= 600 {
		return &AppError{
			code:    http.StatusInternalServerError,
			message: "errs: status code is invalid: " + message,
		}
	}

	return &AppError{code: code, message: message}
}

// error interfaceの実装
func (e AppError) Error() string {
	return e.message
}

func (e AppError) Code() int {
	return e.code
}
