package value

import "errors"

const (
	UnapprovedStatusId = iota
	ApprovedStatusId
)

const (
	UnapprovedStatusName = "未承認"
	ApprovedStatusName   = "承認済"
)

type ClubStatus struct {
	Id   int
	Name string
}

func NewClubStatus(id int) (*ClubStatus, error) {
	if id == UnapprovedStatusId {
		return &ClubStatus{id, UnapprovedStatusName}, nil
	}

	if id == ApprovedStatusId {
		return &ClubStatus{id, ApprovedStatusName}, nil
	}

	return nil, errors.New("status id is invalid")
}

func NewApprovedStatus() *ClubStatus {
	return &ClubStatus{ApprovedStatusId, ApprovedStatusName}
}

func NewUnapprovedStatus() *ClubStatus {
	return &ClubStatus{ApprovedStatusId, UnapprovedStatusName}
}
