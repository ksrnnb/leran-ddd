package response

type GetTeachersResponse struct {
	code int
	body interface{}
}

func NewGetTeachersResponse(code int, body interface{}) *GetTeachersResponse {
	return &GetTeachersResponse{code: code, body: body}
}

func (res GetTeachersResponse) Code() int {
	return res.code
}

func (res GetTeachersResponse) Body() interface{} {
	return res.body
}
