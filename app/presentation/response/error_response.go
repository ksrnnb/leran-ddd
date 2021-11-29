package response

type ErrorResponse struct {
	code    int
	message string
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{code: code, message: message}
}

func (res ErrorResponse) Code() int {
	return res.code
}

func (res ErrorResponse) Body() interface{} {
	return map[string]interface{}{
		"code":    res.code,
		"message": res.message,
	}
}
