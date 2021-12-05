package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/domain/value"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubMockRepository struct {
	clubs []*entity.Club
}

func NewClubMockRepository() repository.ClubRepositoryInterface {
	var clubs []*entity.Club
	for i := 1; i < 3; i++ {
		club, err := entity.NewClub(uint(i), "mockClub"+strconv.Itoa(i), 0)

		if err != nil {
			return nil
		}

		clubs = append(clubs, club)
	}

	return &ClubMockRepository{clubs}
}

func (r ClubMockRepository) GetClubs() ([]*entity.Club, errs.AppErrorInterface) {
	return r.clubs, nil
}

func (r *ClubMockRepository) CreateClub(club *entity.Club) (*entity.Club, errs.AppErrorInterface) {
	id, err := value.NewId(uint(len(r.clubs) + 1))

	if err != nil {
		return nil, errs.NewAppError(http.StatusUnprocessableEntity, err.Error())
	}
	club.Id = *id

	r.clubs = append(r.clubs, club)
	return &entity.Club{}, nil
}

func (r ClubMockRepository) FindByName(name string) (*entity.Club, errs.AppErrorInterface) {
	for _, club := range r.clubs {
		if club.Name.Value == name {
			return club, nil
		}
	}

	return &entity.Club{}, nil
}
