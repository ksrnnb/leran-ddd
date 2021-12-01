package usecase

import (
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
	"github.com/ksrnnb/learn-ddd/usecase/output"
)

type GetClubsUsecase struct {
	repo repository.ClubRepositoryInterface
}

func NewGetClubsUsecase(repo repository.ClubRepositoryInterface) *GetClubsUsecase {
	return &GetClubsUsecase{repo: repo}
}

func (u GetClubsUsecase) GetClubs() (*output.GetClubsOutput, errs.AppErrorInterface) {
	domainClubs, err := u.repo.GetClubs()

	if err != nil {
		return nil, err
	}

	return output.NewGetClubsResponse(domainClubs), nil
}
