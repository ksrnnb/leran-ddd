package infrastructure

import "github.com/ksrnnb/learn-ddd/domain/entity"

type TeacherMockRepository struct{}

var teacherMock []*entity.Teacher

func NewTeacherMockRepository(id uint, name string) *TeacherRepository {
	t, err := entity.NewTeacher(id, name)

	if err != nil {
		return nil
	}

	teacherMock = append(teacherMock, t)
	return &TeacherRepository{}
}

func (r TeacherMockRepository) GetTeachers() ([]*entity.Teacher, error) {
	return teacherMock, nil
}
