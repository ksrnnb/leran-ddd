package infrastructure

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type TeacherRepository struct{}

func NewTeacherRepository() *TeacherRepository {
	return &TeacherRepository{}
}

func (r TeacherRepository) GetTeachers() ([]*entity.Teacher, errs.AppErrorInterface) {
	// TODO: 実装
	return make([]*entity.Teacher, 0), nil
}
