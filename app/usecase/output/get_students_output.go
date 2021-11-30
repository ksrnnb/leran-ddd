package output

import "github.com/ksrnnb/learn-ddd/domain/entity"

type GetStudentsOutput struct {
	Students []*Student
}

func NewGetStudentsResponse(domainStudents []*entity.Student) *GetStudentsOutput {
	students := make([]*Student, 0)

	for _, domainStudent := range domainStudents {
		newStudent := &Student{
			Id:   domainStudent.Id.Value,
			Name: domainStudent.Name.Value,
		}

		students = append(students, newStudent)
	}

	return &GetStudentsOutput{Students: students}
}

type Student struct {
	Id   uint
	Name string
}
