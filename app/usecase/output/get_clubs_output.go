package output

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
)

type GetClubsOutput struct {
	Clubs []*Club
}

func NewGetClubsResponse(domainClubs []*entity.Club) *GetClubsOutput {
	clubs := make([]*Club, 0)

	for _, domainClub := range domainClubs {
		newClub := &Club{
			Id:     domainClub.Id.Value,
			Name:   domainClub.Name.Value,
			Status: domainClub.Status.Name,
		}

		clubs = append(clubs, newClub)
	}

	return &GetClubsOutput{Clubs: clubs}
}

type Club struct {
	Id     uint
	Name   string
	Status string
}
