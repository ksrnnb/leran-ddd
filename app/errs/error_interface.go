package errs

type AppErrorInterface interface {
	Error() string
	Code() int
}
