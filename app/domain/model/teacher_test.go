package model

import "testing"

func TestNewTeacherTooShortName(t *testing.T) {
	tooShortName := "a"
	_, err := NewTeacher(tooShortName)

	if err == nil {
		t.Errorf("error while testing too short name")
	}

	borderShortName := "ab"
	_, err = NewTeacher(borderShortName)

	if err != nil {
		t.Errorf("error while testing border short name")
	}
}

func TestNewTeacherTooLongName(t *testing.T) {
	tooLongName := "123456789012345678901"
	_, err := NewTeacher(tooLongName)

	if err == nil {
		t.Errorf("error while testing too long name")
	}

	borderLongName := "12345678901234567890"
	_, err = NewTeacher(borderLongName)

	if err != nil {
		t.Error("error while testing border long name")
	}
}
