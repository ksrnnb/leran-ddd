package infrastructure

import (
	"strconv"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
)

type StudentMockRepository struct {
	students []*entity.Student
}

func NewStudentMockRepository() repository.StudentRepositoryInterface {
	var students []*entity.Student
	for i := 1; i < 10; i++ {
		student, err := entity.NewStudent(uint(i), "mockStudent"+strconv.Itoa(i))

		if err != nil {
			return nil
		}

		students = append(students, student)
	}

	return &StudentMockRepository{students}
}

func (r StudentMockRepository) GetStudents() ([]*entity.Student, errs.AppErrorInterface) {
	return r.students, nil
}

func (r StudentMockRepository) Find(studentId int) (*entity.Student, errs.AppErrorInterface) {
	for _, t := range r.students {
		if t.Id.Value == uint(studentId) {
			return t, nil
		}
	}

	return &entity.Student{}, nil
}
