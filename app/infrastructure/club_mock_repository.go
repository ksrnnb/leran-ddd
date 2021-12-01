package infrastructure

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubMockRepository struct{}

var clubMock []*entity.Club

func NewClubMockRepository(id uint, name string, statusId int) repository.ClubRepositoryInterface {
	s, err := entity.NewClub(id, name, statusId)

	if err != nil {
		return nil
	}

	clubMock = append(clubMock, s)
	return &ClubMockRepository{}
}

func (r ClubMockRepository) GetClubs() ([]*entity.Club, errs.AppErrorInterface) {
	return clubMock, nil
}