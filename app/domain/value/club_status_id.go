package value

import "errors"

const (
	UnapprovedStatusId = iota
	ApprovedStatusId
)

type ClubStatus struct {
	Id   int
	Name string
}

func NewClubStatus(id int) (*ClubStatus, error) {
	if id == UnapprovedStatusId {
		return &ClubStatus{id, "未承認"}, nil
	}

	if id == ApprovedStatusId {
		return &ClubStatus{id, "承認済"}, nil
	}

	return nil, errors.New("status id is invalid")
}
