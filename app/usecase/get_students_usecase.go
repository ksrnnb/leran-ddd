package usecase

import (
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/errs"
	"github.com/ksrnnb/learn-ddd/usecase/output"
)

type GetStudentsUsecase struct {
	repo repository.StudentRepositoryInterface
}

func NewGetStudentsUsecase(repo repository.StudentRepositoryInterface) *GetStudentsUsecase {
	return &GetStudentsUsecase{repo: repo}
}

func (u GetStudentsUsecase) GetStudents() (*output.GetStudentsOutput, errs.AppErrorInterface) {
	domainStudents, err := u.repo.GetStudents()

	if err != nil {
		return nil, err
	}

	return output.NewGetStudentsResponse(domainStudents), nil
}
