package response

type CreateClubResponse struct {
	code int
	body interface{}
}

func NewCreateClubResponse(code int, body interface{}) *CreateClubResponse {
	return &CreateClubResponse{code: code, body: body}
}

func (res CreateClubResponse) Code() int {
	return res.code
}

func (res CreateClubResponse) Body() interface{} {
	return res.body
}
