package value

import "errors"

const (
	nameUpperLimit = 20
	nameLowerLimit = 2
)

type Name struct {
	Value string
}

func NewName(name string) (*Name, error) {
	if len(name) > nameUpperLimit {
		return nil, errors.New("name is too long")
	}

	if (len(name)) < nameLowerLimit {
		return nil, errors.New("name is too short")
	}

	return &Name{Value: name}, nil
}
