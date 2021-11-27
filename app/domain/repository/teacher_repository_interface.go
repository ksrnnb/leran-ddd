package repository

import "github.com/ksrnnb/learn-ddd/domain/model"

type TeacherRepositoryInterface interface {
	GetTeachers() []*model.Teacher
}
