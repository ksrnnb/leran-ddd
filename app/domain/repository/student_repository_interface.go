package repository

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type StudentRepositoryInterface interface {
	GetStudents() ([]*entity.Student, errs.AppErrorInterface)
	Find(studentId int) (*entity.Student, errs.AppErrorInterface)
}
