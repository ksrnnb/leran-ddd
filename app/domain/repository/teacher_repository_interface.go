package repository

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type TeacherRepositoryInterface interface {
	GetTeachers() ([]*entity.Teacher, errs.AppErrorInterface)
}
