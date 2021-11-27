package infrastructure

import "github.com/ksrnnb/learn-ddd/domain/entity"

type TeacherRepository struct{}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{}
}

func (r TeacherRepository) GetTeachers() ([]*entity.Teacher, error) {
	// TODO: 実装
	return make([]*entity.Teacher, 0), nil
}
