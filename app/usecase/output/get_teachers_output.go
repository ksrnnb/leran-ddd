package output

import "github.com/ksrnnb/learn-ddd/domain/entity"

type GetTeachersOutput struct {
	Teachers []*Teacher
}

func NewGetTeachersResponse(domainTeachers []*entity.Teacher) *GetTeachersOutput {
	teachers := make([]*Teacher, 0)

	for _, domainTeacher := range domainTeachers {
		newTeacher := &Teacher{
			Id:   domainTeacher.Id.Value,
			Name: domainTeacher.Name.Value,
		}

		teachers = append(teachers, newTeacher)
	}

	return &GetTeachersOutput{Teachers: teachers}
}

type Teacher struct {
	Id   uint
	Name string
}
