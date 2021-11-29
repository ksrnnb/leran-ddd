package entity

import "testing"

func TestNewTeacherTooShortName(t *testing.T) {
	tooShortName := "a"
	_, err := NewTeacher(1, tooShortName)

	if err == nil {
		t.Errorf("error while testing too short name")
	}

	borderShortName := "ab"
	_, err = NewTeacher(1, borderShortName)

	if err != nil {
		t.Errorf("error while testing border short name")
	}
}

func TestNewTeacherTooLongName(t *testing.T) {
	tooLongName := "123456789012345678901"
	_, err := NewTeacher(1, tooLongName)

	if err == nil {
		t.Errorf("error while testing too long name")
	}

	borderLongName := "12345678901234567890"
	_, err = NewTeacher(1, borderLongName)

	if err != nil {
		t.Error("error while testing border long name")
	}
}

func TestNewTeacherIdIsZero(t *testing.T) {
	_, err := NewTeacher(0, "test")

	if err == nil {
		t.Errorf("error while testing teacher id is zero")
	}
}
