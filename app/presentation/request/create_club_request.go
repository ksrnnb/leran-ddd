package request

type CreateClubRequest struct {
	Name       string `json:"name" validate:"required"`
	StudentIds []int  `json:"studentIds" validate:"required"`
	TeacherId  int    `json:"teacherId" validate:"required"`
}
