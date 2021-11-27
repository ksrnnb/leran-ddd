package entity

import "github.com/ksrnnb/learn-ddd/domain/value"

type Teacher struct {
	Id   value.Id
	Name value.Name
}

// 教員の生成
func NewTeacher(id uint, name string) (*Teacher, error) {
	nameObject, err := value.NewName(name)

	if err != nil {
		return nil, err
	}

	idObject, err := value.NewId(id)

	if err != nil {
		return nil, err
	}

	return &Teacher{
		Id:   *idObject,
		Name: *nameObject,
	}, nil
}
