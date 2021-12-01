package entity

import (
	"testing"

	"github.com/ksrnnb/learn-ddd/domain/value"
)

func TestNewClubTooShortName(t *testing.T) {
	tooShortName := ""

	_, err := NewClub(1, tooShortName, value.UnapprovedStatusId)

	if err == nil {
		t.Error("error while creating new club too short name")
	}

	borderTooShortName := "a"

	_, err = NewClub(1, borderTooShortName, value.UnapprovedStatusId)

	if err != nil {
		t.Error("error while creating new club border short name")
	}
}

func TestNewClubTooLongName(t *testing.T) {
	tooLongName := "abcdeabcdeabdceabcdef"

	_, err := NewClub(1, tooLongName, value.UnapprovedStatusId)

	if err == nil {
		t.Error("error while creating new club too long name")
	}

	borderTooLongName := "abcdeabcdeabdceabcde"

	_, err = NewClub(1, borderTooLongName, value.UnapprovedStatusId)

	if err != nil {
		t.Error("error while creating new club border long name")
	}
}

func TestNewClubIdIsZero(t *testing.T) {
	_, err := NewClub(0, "club", value.UnapprovedStatusId)

	if err == nil {
		t.Error("error while creating new club id is zero")
	}
}

func TestNewClubInvalidStatusId(t *testing.T) {
	_, err := NewClub(1, "club", -1)

	if err == nil {
		t.Error("error while creating new club status id is minus")
	}

	_, err = NewClub(1, "club", 2)

	if err == nil {
		t.Error("error while creating new club status id is over 2")
	}
}
