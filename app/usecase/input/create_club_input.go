package input

type CreateClubInput struct {
	Name       string
	StudentIds []int
	TeacherId  int
}

func NewCreateClubInput(name string, studentIds []int, teacherId int) *CreateClubInput {
	return &CreateClubInput{
		Name:       name,
		StudentIds: studentIds,
		TeacherId:  teacherId,
	}
}
