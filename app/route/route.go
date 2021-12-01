package route

import (
	"github.com/ksrnnb/learn-ddd/container"
	"github.com/ksrnnb/learn-ddd/presentation/request"
	"github.com/labstack/echo/v4"
)

var c *container.Container

func RegisterRoute(e *echo.Echo) {
	e.GET("/teachers", getTeachers)
	e.GET("/students", getStudents)
	e.GET("/clubs", getClubs)
}

func getTeachers(context echo.Context) error {
	c = container.NewContainer()
	req := request.NewGetTeachersRequest()
	res := c.TeacherController.GetTeachers(req)
	return context.JSON(res.Code(), res.Body())
}

func getStudents(context echo.Context) error {
	c = container.NewContainer()
	req := request.NewGetStudentsRequest()
	res := c.StudentController.GetStudents(req)
	return context.JSON(res.Code(), res.Body())
}

func getClubs(context echo.Context) error {
	c = container.NewContainer()
	req := request.NewGetClubsRequest()
	res := c.ClubController.GetClubs(req)
	return context.JSON(res.Code(), res.Body())
}
