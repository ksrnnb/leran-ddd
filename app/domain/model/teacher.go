package model

import "github.com/ksrnnb/learn-ddd/domain/value"

type Teacher struct {
	Id   int
	Name value.Name
}

// 教員の生成
func NewTeacher(name string) (*Teacher, error) {
	nameObject, err := value.NewName(name)

	if err != nil {
		return nil, err
	}

	return &Teacher{Name: *nameObject}, nil
}
