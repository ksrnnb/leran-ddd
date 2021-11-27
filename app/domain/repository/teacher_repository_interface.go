package repository

import "github.com/ksrnnb/learn-ddd/domain/entity"

type TeacherRepositoryInterface interface {
	GetTeachers() ([]*entity.Teacher, error)
}
