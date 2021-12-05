package output

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
)

type CreateClubOutput struct {
	Club *Club
}

func NewCreateClubOutput(domainClub *entity.Club) *CreateClubOutput {
	newClub := &Club{
		Id:     domainClub.Id.Value,
		Name:   domainClub.Name.Value,
		Status: domainClub.Status.Name,
	}

	return &CreateClubOutput{Club: newClub}
}
