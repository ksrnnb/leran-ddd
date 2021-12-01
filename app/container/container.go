package container

import (
	"github.com/ksrnnb/learn-ddd/infrastructure"
	"github.com/ksrnnb/learn-ddd/infrastructure/db"
	"github.com/ksrnnb/learn-ddd/presentation"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type Container struct {
	TeacherController *presentation.TeacherController
	StudentController *presentation.StudentController
	ClubController    *presentation.ClubController
}

// TODO: もう少しいい方法を調べる
func NewContainer() *Container {
	db, err := db.NewDBClient()
	if err != nil {
		panic(err)
	}

	container := &Container{}

	u := usecase.NewGetTeachersUsecase(infrastructure.NewTeacherRepository(db))
	container.TeacherController = presentation.NewTeacherController(*u)

	su := usecase.NewGetStudentsUsecase(infrastructure.NewStudentRepository(db))
	container.StudentController = presentation.NewStudentController(*su)

	cu := usecase.NewGetClubsUsecase(infrastructure.NewClubRepository(db))
	container.ClubController = presentation.NewClubController(*cu)

	return container
}
