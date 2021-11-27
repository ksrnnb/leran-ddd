package route

import (
	"github.com/ksrnnb/learn-ddd/infrastructure"
	"github.com/ksrnnb/learn-ddd/presentation"
	"github.com/ksrnnb/learn-ddd/presentation/response"
	"github.com/ksrnnb/learn-ddd/usecase"
	"github.com/labstack/echo/v4"
)

var teacherController *presentation.TeacherController

func init() {
	// TODO: もう少し良い渡し方がないか検討
	// di container...
	u := usecase.NewGetTeachersUsecase(infrastructure.NewTeacherRepository())
	teacherController = presentation.NewTeacherController(*u)
}

func RegisterRoute(e *echo.Echo) {
	e.GET("/teachers", echoWrapper(teacherController.GetTeachers()))
}

func echoWrapper(res response.Response) func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.JSON(res.Code(), res.Body())
	}
}
