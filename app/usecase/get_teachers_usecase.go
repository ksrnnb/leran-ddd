package usecase

import (
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/usecase/output"
)

type GetTeachersUsecase struct {
	repo repository.TeacherRepositoryInterface
}

func NewGetTeachersUsecase(repo repository.TeacherRepositoryInterface) *GetTeachersUsecase {
	return &GetTeachersUsecase{repo: repo}
}

func (u GetTeachersUsecase) GetTeachers() (*output.GetTeachersOutput, error) {
	domainTeachers, err := u.repo.GetTeachers()

	if err != nil {
		return nil, err
	}

	return output.NewGetTeachersResponse(domainTeachers), nil
}
