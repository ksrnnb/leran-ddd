package presentation

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/logger"
	"github.com/ksrnnb/learn-ddd/presentation/request"
	"github.com/ksrnnb/learn-ddd/presentation/response"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type TeacherController struct {
	getTeachersUsecase usecase.GetTeachersUsecase
}

func NewTeacherController(getTeachersUsercase usecase.GetTeachersUsecase) *TeacherController {
	return &TeacherController{getTeachersUsecase: getTeachersUsercase}
}

func (c TeacherController) GetTeachers(req *request.GetTeachersRequest) response.Response {
	out, err := c.getTeachersUsecase.GetTeachers()

	if err != nil {
		logger.Println(err.Error())
		return response.NewErrorResponse(err.Code(), err.Error())
	}

	res := response.NewGetTeachersResponse(http.StatusOK, out.Teachers)
	return res
}
