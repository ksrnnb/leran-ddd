package usecase

import (
	"os"
	"testing"

	"github.com/ksrnnb/learn-ddd/infrastructure"
)

var getTeachersUsecase *GetTeachersUsecase
var getStudentsUsecase *GetStudentsUsecase

func TestMain(m *testing.M) {
	setUp()
	v := m.Run()
	os.Exit(v)
}

func setUp() {
	mockTeacherRepo := infrastructure.NewTeacherMockRepository(1, "テスト")
	getTeachersUsecase = NewGetTeachersUsecase(mockTeacherRepo)

	mockStudentRepo := infrastructure.NewStudentMockRepository(1, "テスト")
	getStudentsUsecase = NewGetStudentsUsecase(mockStudentRepo)
}
