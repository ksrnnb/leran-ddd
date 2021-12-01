package presentation

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/logger"
	"github.com/ksrnnb/learn-ddd/presentation/request"
	"github.com/ksrnnb/learn-ddd/presentation/response"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type ClubController struct {
	getClubsUsecase usecase.GetClubsUsecase
}

func NewClubController(getClubsUsercase usecase.GetClubsUsecase) *ClubController {
	return &ClubController{getClubsUsecase: getClubsUsercase}
}

func (c ClubController) GetClubs(req *request.GetClubsRequest) response.Response {
	out, err := c.getClubsUsecase.GetClubs()

	if err != nil {
		logger.Println(err.Error())
		return response.NewErrorResponse(err.Code(), err.Error())
	}

	res := response.NewGetClubsResponse(http.StatusOK, out.Clubs)
	return res
}
