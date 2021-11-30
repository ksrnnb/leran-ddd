package usecase

import (
	"testing"
)

func TestGetStudents(t *testing.T) {
	out, err := getStudentsUsecase.GetStudents()

	if err != nil {
		t.Error("error while testing get students length")
		return
	}

	students := out.Students

	if len(students) == 0 {
		t.Error("error while testing get students length")
		return
	}

	student := students[0]

	if student.Id != 1 {
		t.Error("error while testing get student's id")
	}

	if student.Name != "テスト" {
		t.Error("error while testing get student's name")
	}
}
