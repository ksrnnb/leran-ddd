package presentation

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/logger"
	"github.com/ksrnnb/learn-ddd/presentation/request"
	"github.com/ksrnnb/learn-ddd/presentation/response"
	"github.com/ksrnnb/learn-ddd/usecase"
	"github.com/ksrnnb/learn-ddd/usecase/input"
)

type ClubController struct {
	getClubsUsecase   usecase.GetClubsUsecase
	createClubUsecase usecase.CreateClubUsecase
}

func NewClubController(
	getClubsUsercase usecase.GetClubsUsecase,
	createClubUsecase usecase.CreateClubUsecase,
) *ClubController {
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

func (c ClubController) CreateClub(req *request.CreateClubRequest) response.Response {
	in := input.NewCreateClubInput(req.Name, req.StudentIds, req.TeacherId)
	out, err := c.createClubUsecase.CreateClub(in)

	if err != nil {
		logger.Println(err.Error())
		return response.NewErrorResponse(err.Code(), err.Error())
	}

	res := response.NewCreateClubResponse(http.StatusCreated, out.Club)
	return res
}
