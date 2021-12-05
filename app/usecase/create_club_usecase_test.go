package usecase

import (
	"testing"

	"github.com/ksrnnb/learn-ddd/domain/value"
	"github.com/ksrnnb/learn-ddd/usecase/input"
)

func TestCreateClubInvalidName(t *testing.T) {
	in := input.NewCreateClubInput("", []int{0, 1, 2}, 0)
	_, err := createClubUsecase.CreateClub(in)

	if err == nil {
		t.Errorf("error while testing create club invalid name")
	}
}

func TestCreateClubInvalidStudentIds(t *testing.T) {
	in := input.NewCreateClubInput("testClub", []int{0, 1, 2}, 1)
	_, err := createClubUsecase.CreateClub(in)

	if err == nil {
		t.Errorf("error while testing create club invalid student ids")
	}
}

func TestCreateClubInvalidTeacherId(t *testing.T) {
	in := input.NewCreateClubInput("testClub", []int{0, 1, 2}, 0)
	_, err := createClubUsecase.CreateClub(in)

	if err == nil {
		t.Errorf("error while testing create club invalid teacher id")
	}
}

func TestCreateClubDuplicatedName(t *testing.T) {
	clubName := "testClub1"
	in := input.NewCreateClubInput(clubName, []int{1, 2, 3}, 1)
	_, err := createClubUsecase.CreateClub(in)

	if err != nil {
		t.Errorf("error while testing create club valid input")
	}

	in = input.NewCreateClubInput(clubName, []int{1, 2, 3}, 1)
	_, err = createClubUsecase.CreateClub(in)

	if err == nil {
		t.Errorf("error while testing create club duplicated name")
	}
}

func TestCreateClubStatus(t *testing.T) {
	in := input.NewCreateClubInput("testClub2", []int{1, 2, 3, 4}, 1)
	_, err := createClubUsecase.CreateClub(in)

	if err != nil {
		t.Errorf("error while testing create club valid input")
	}

	out, err := getClubsUsecase.GetClubs()

	if err != nil {
		t.Errorf("error while testing create club getting clubs")
	}

	lastestClub := out.Clubs[len(out.Clubs)-1]
	if lastestClub.Status != value.UnapprovedStatusName {
		t.Errorf("error while testing create club unapproved status name")
	}

	in = input.NewCreateClubInput("testClub3", []int{1, 2, 3, 4, 5}, 1)
	_, err = createClubUsecase.CreateClub(in)

	out, err = getClubsUsecase.GetClubs()

	if err != nil {
		t.Errorf("error while testing create club getting clubs")
	}

	lastestClub = out.Clubs[len(out.Clubs)-1]
	if lastestClub.Status != value.ApprovedStatusName {
		t.Errorf("error while testing create club approved status name")
	}
}
