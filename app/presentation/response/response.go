package response

type Response interface {
	Code() int
	Body() interface{}
}
