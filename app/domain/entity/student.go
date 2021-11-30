package entity

import "github.com/ksrnnb/learn-ddd/domain/value"

type Student struct {
	Id   value.Id
	Name value.Name
}

func NewStudent(id uint, name string) (*Student, error) {
	idObject, err := value.NewId(id)

	if err != nil {
		return nil, err
	}

	NameObject, err := value.NewName(name)

	if err != nil {
		return nil, err
	}

	return &Student{Id: *idObject, Name: *NameObject}, nil
}
