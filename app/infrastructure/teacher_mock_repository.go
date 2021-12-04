package infrastructure

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
)

type TeacherMockRepository struct{}

var teacherMock []*entity.Teacher

func NewTeacherMockRepository(id uint, name string) repository.TeacherRepositoryInterface {
	t, err := entity.NewTeacher(id, name)

	if err != nil {
		return nil
	}

	teacherMock = append(teacherMock, t)
	return &TeacherMockRepository{}
}

func (r TeacherMockRepository) GetTeachers() ([]*entity.Teacher, errs.AppErrorInterface) {
	return teacherMock, nil
}

func (r TeacherMockRepository) Find(teacherId int) (*entity.Teacher, errs.AppErrorInterface) {
	return &entity.Teacher{}, nil
}
