package service

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
)

type ClubService struct {
	repo repository.ClubRepositoryInterface
}

func NewClubService(repo repository.ClubRepositoryInterface) *ClubService {
	return &ClubService{repo: repo}
}

func (s ClubService) IsDuplicatedName(club *entity.Club) (bool, error) {
	c, err := s.repo.FindByName(club.Name.Value)

	if err != nil {
		return false, err
	}

	return c == nil, nil
}
