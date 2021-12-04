package container

import (
	"github.com/ksrnnb/learn-ddd/domain/factory"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/infrastructure"
	"github.com/ksrnnb/learn-ddd/infrastructure/db"
	"github.com/ksrnnb/learn-ddd/presentation"
	"github.com/ksrnnb/learn-ddd/usecase"
)

type Container struct {
	TeacherController *presentation.TeacherController
	StudentController *presentation.StudentController
	ClubController    *presentation.ClubController
	TeacherRepository repository.TeacherRepositoryInterface
	StudentRepository repository.StudentRepositoryInterface
	ClubRepository    repository.ClubRepositoryInterface
}

// TODO: もう少しいい方法を調べる
func NewContainer() *Container {

	container := &Container{}
	container.addRepository()

	u := usecase.NewGetTeachersUsecase(container.TeacherRepository)
	su := usecase.NewGetStudentsUsecase(container.StudentRepository)
	gcu := usecase.NewGetClubsUsecase(container.ClubRepository)

	cf := factory.NewClubFactory(container.ClubRepository, container.StudentRepository, container.TeacherRepository)
	ccu := usecase.NewCreateClubUsecase(container.ClubRepository, *cf)

	container.TeacherController = presentation.NewTeacherController(*u)
	container.StudentController = presentation.NewStudentController(*su)
	container.ClubController = presentation.NewClubController(*gcu, *ccu)

	return container
}

func (c *Container) addRepository() {
	db, err := db.NewDBClient()
	if err != nil {
		panic(err)
	}

	c.TeacherRepository = infrastructure.NewTeacherRepository(db)
	c.StudentRepository = infrastructure.NewStudentRepository(db)
	c.ClubRepository = infrastructure.NewClubRepository(db)
}
