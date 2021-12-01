package entity

import (
	"github.com/ksrnnb/learn-ddd/domain/value"
)

type Club struct {
	Id     value.Id
	Name   value.ClubName
	Status value.ClubStatus
}

func NewClub(id uint, name string, statusId int) (*Club, error) {
	idObject, err := value.NewId(id)

	if err != nil {
		return nil, err
	}

	nameObject, err := value.NewClubName(name)

	if err != nil {
		return nil, err
	}

	statusObject, err := value.NewClubStatus(statusId)

	if err != nil {
		return nil, err
	}

	return &Club{*idObject, *nameObject, *statusObject}, nil
}
