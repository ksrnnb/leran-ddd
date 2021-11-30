package usecase

import (
	"testing"
)

func TestGetTeachers(t *testing.T) {
	out, err := getTeachersUsecase.GetTeachers()

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
