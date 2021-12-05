package infrastructure

import (
	"strconv"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
)

type TeacherMockRepository struct {
	teachers []*entity.Teacher
}

func NewTeacherMockRepository() repository.TeacherRepositoryInterface {
	var teachers []*entity.Teacher
	for i := 1; i < 3; i++ {
		teacher, err := entity.NewTeacher(uint(i), "mockTeacher"+strconv.Itoa(i))

		if err != nil {
			return nil
		}

		teachers = append(teachers, teacher)
	}

	return &TeacherMockRepository{teachers}
}

func (r TeacherMockRepository) GetTeachers() ([]*entity.Teacher, errs.AppErrorInterface) {
	return r.teachers, nil
}

func (r TeacherMockRepository) Find(teacherId int) (*entity.Teacher, errs.AppErrorInterface) {
	for _, t := range r.teachers {
		if t.Id.Value == uint(teacherId) {
			return t, nil
		}
	}

	return &entity.Teacher{}, nil
}
