package response

type GetClubsResponse struct {
	code int
	body interface{}
}

func NewGetClubsResponse(code int, body interface{}) *GetClubsResponse {
	return &GetClubsResponse{code: code, body: body}
}

func (res GetClubsResponse) Code() int {
	return res.code
}

func (res GetClubsResponse) Body() interface{} {
	return res.body
}
