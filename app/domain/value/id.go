package value

import "errors"

type Id struct {
	Value uint
}

func NewId(id uint) (*Id, error) {
	if id == 0 {
		return nil, errors.New("id should be positive integer")
	}

	return &Id{Value: id}, nil
}
