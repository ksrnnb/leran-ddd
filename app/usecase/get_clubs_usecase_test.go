package usecase

import (
	"testing"
)

func TestGetClubs(t *testing.T) {
	out, err := getClubsUsecase.GetClubs()

	if err != nil {
		t.Error("error while testing get clubs length")
		return
	}

	clubs := out.Clubs

	if len(clubs) == 0 {
		t.Error("error while testing get clubs length")
		return
	}

	club := clubs[0]

	if club.Id == 0 {
		t.Error("error while testing get club's id")
	}
}
