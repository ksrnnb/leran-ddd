package presentation

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/logger"
	"github.com/ksrnnb/learn-ddd/presentation/request"
	"github.com/ksrnnb/learn-ddd/presentation/response"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type StudentController struct {
	getStudentsUsecase usecase.GetStudentsUsecase
}

func NewStudentController(getStudentsUsercase usecase.GetStudentsUsecase) *StudentController {
	return &StudentController{getStudentsUsecase: getStudentsUsercase}
}

func (c StudentController) GetStudents(req *request.GetStudentsRequest) response.Response {
	out, err := c.getStudentsUsecase.GetStudents()

	if err != nil {
		logger.Println(err.Error())
		return response.NewErrorResponse(err.Code(), err.Error())
	}

	res := response.NewGetStudentsResponse(http.StatusOK, out.Students)
	return res
}
