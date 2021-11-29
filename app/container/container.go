package container

import (
	"github.com/ksrnnb/learn-ddd/infrastructure"
	"github.com/ksrnnb/learn-ddd/infrastructure/db"
	"github.com/ksrnnb/learn-ddd/presentation"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type Container struct {
	TeacherController *presentation.TeacherController
}

func NewContainer() *Container {
	db, err := db.NewDBClient()
	if err != nil {
		panic(err)
	}

	container := &Container{}

	u := usecase.NewGetTeachersUsecase(infrastructure.NewTeacherRepository(db))
	container.TeacherController = presentation.NewTeacherController(*u)

	return container
}
