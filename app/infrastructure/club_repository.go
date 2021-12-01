package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubRepository struct {
	db *sql.DB
}

// レポジトリの生成
func NewClubRepository(db *sql.DB) *ClubRepository {
	return &ClubRepository{db: db}
}

func (r ClubRepository) GetClubs() ([]*entity.Club, errs.AppErrorInterface) {
	rows, err := r.db.Query("SELECT id, name, status_id from clubs")

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	var clubs []*entity.Club
	for rows.Next() {
		club, err := r.scanClub(rows)

		if err != nil {
			return nil, err
		}

		clubs = append(clubs, club)
	}

	err = rows.Err()

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return clubs, nil
}

func (r ClubRepository) scanClub(rows *sql.Rows) (*entity.Club, errs.AppErrorInterface) {
	var id uint
	var name string
	var statusId int
	err := rows.Scan(&id, &name, &statusId)
	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	club, err := entity.NewClub(id, name, statusId)

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return club, nil
}
