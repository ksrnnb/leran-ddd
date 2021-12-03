package usecase

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/factory"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/domain/service"
	"github.com/ksrnnb/learn-ddd/errs"
	"github.com/ksrnnb/learn-ddd/usecase/input"
	"github.com/ksrnnb/learn-ddd/usecase/output"
)

type CreateClubUsecase struct {
	repo    repository.ClubRepositoryInterface
	factory factory.ClubFactory
}

func NewCreateClubUsecase(repo repository.ClubRepositoryInterface) *CreateClubUsecase {
	return &CreateClubUsecase{repo: repo}
}

func (u CreateClubUsecase) CreateClub(in *input.CreateClubInput) (*output.CreateClubOutput, errs.AppErrorInterface) {
	club, appErr := u.factory.NewClubWithStudentsAndTeacher(in.Name, in.StudentIds, in.TeacherId)

	if appErr != nil {
		return nil, appErr
	}

	clubService := service.NewClubService(u.repo)

	isDuplicated, err := clubService.IsDuplicatedName(club)

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	if isDuplicated {
		return nil, errs.NewAppError(http.StatusUnprocessableEntity, "usecase: club name is duplicated")
	}

	appErr = u.repo.CreateClub(club)

	if appErr != nil {
		return nil, appErr
	}

	return output.NewCreateClubOutput(club), nil
}
