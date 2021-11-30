package response

type GetStudentsResponse struct {
	code int
	body interface{}
}

func NewGetStudentsResponse(code int, body interface{}) *GetStudentsResponse {
	return &GetStudentsResponse{code: code, body: body}
}

func (res GetStudentsResponse) Code() int {
	return res.code
}

func (res GetStudentsResponse) Body() interface{} {
	return res.body
}
