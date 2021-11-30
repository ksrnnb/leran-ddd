package infrastructure

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
)

type StudentMockRepository struct{}

var studentMock []*entity.Student

func NewStudentMockRepository(id uint, name string) repository.StudentRepositoryInterface {
	s, err := entity.NewStudent(id, name)

	if err != nil {
		return nil
	}

	studentMock = append(studentMock, s)
	return &StudentMockRepository{}
}

func (r StudentMockRepository) GetStudents() ([]*entity.Student, errs.AppErrorInterface) {
	return studentMock, nil
}
