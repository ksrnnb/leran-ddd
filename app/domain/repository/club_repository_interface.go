package repository

import (
	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubRepositoryInterface interface {
	GetClubs() ([]*entity.Club, errs.AppErrorInterface)
	CreateClub(*entity.Club) errs.AppErrorInterface
	FindByName(name string) (*entity.Club, errs.AppErrorInterface)
}
