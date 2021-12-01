package value

import "errors"

const (
	clubNameUpperLimit = 20
	clubNameLowerLimit = 1
)

type ClubName struct {
	Value string
}

func NewClubName(name string) (*ClubName, error) {
	if len(name) < clubNameLowerLimit {
		return nil, errors.New("club name is too short")
	}

	if len(name) > clubNameUpperLimit {
		return nil, errors.New("club name is too long")
	}

	return &ClubName{Value: name}, nil
}
