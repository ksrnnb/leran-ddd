package usecase

import (
	"os"
	"testing"

	"github.com/ksrnnb/learn-ddd/domain/factory"
	"github.com/ksrnnb/learn-ddd/infrastructure"
)

var getTeachersUsecase *GetTeachersUsecase
var getStudentsUsecase *GetStudentsUsecase
var getClubsUsecase *GetClubsUsecase
var createClubUsecase *CreateClubUsecase

func TestMain(m *testing.M) {
	setUp()
	v := m.Run()
	os.Exit(v)
}

func setUp() {
	mockTeacherRepo := infrastructure.NewTeacherMockRepository()
	getTeachersUsecase = NewGetTeachersUsecase(mockTeacherRepo)

	mockStudentRepo := infrastructure.NewStudentMockRepository()
	getStudentsUsecase = NewGetStudentsUsecase(mockStudentRepo)

	mockClubRepo := infrastructure.NewClubMockRepository()
	getClubsUsecase = NewGetClubsUsecase(mockClubRepo)

	mockFactory := factory.NewClubFactory(mockClubRepo, mockStudentRepo, mockTeacherRepo)
	createClubUsecase = NewCreateClubUsecase(mockClubRepo, *mockFactory)
}
