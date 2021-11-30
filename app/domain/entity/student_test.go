package entity

import "testing"

func TestNewStudentTooShortName(t *testing.T) {
	tooShortName := "a"

	_, err := NewStudent(1, tooShortName)

	if err == nil {
		t.Error("error while creating new user too short name")
	}

	borderTooShortName := "ab"

	_, err = NewStudent(1, borderTooShortName)

	if err != nil {
		t.Error("error while creating new user border short name")
	}
}

func TestNewStudentTooLongName(t *testing.T) {
	tooLongName := "abcdeabcdeabdceabcdef"

	_, err := NewStudent(1, tooLongName)

	if err == nil {
		t.Error("error while creating new user too long name")
	}

	borderTooLongName := "abcdeabcdeabdceabcde"

	_, err = NewStudent(1, borderTooLongName)

	if err != nil {
		t.Error("error while creating new user border long name")
	}
}

func TestNewStudentIdIsZero(t *testing.T) {
	_, err := NewStudent(0, "student")

	if err == nil {
		t.Error("error while creating new user id is zero")
	}
}
