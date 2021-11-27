package usecase

import (
	"os"
	"testing"

	"github.com/ksrnnb/learn-ddd/infrastructure"
)

var getUserUsecase *GetTeachersUsecase

func TestMain(m *testing.M) {
	setUp()
	v := m.Run()
	os.Exit(v)
}

func setUp() {
	mockRepo := infrastructure.NewTeacherMockRepository(1, "テスト")
	getUserUsecase = NewGetTeachersUsecase(mockRepo)
}

func TestGetTeachers(t *testing.T) {
	out, err := getUserUsecase.GetTeachers()

	if err != nil {
		t.Error("error while testing get teachers length")
		return
	}

	teachers := out.Teachers

	if len(teachers) == 0 {
		t.Error("error while testing get teachers length")
		return
	}

	teacher := teachers[0]

	if teacher.Id != 1 {
		t.Error("error while testing get teacher's id")
	}

	if teacher.Name != "テスト" {
		t.Error("error while testing get teacher's name")
	}
}
