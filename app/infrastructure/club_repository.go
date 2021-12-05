package infrastructure

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/domain/value"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubRepository struct {
	db *sql.DB
}

// レポジトリの生成
func NewClubRepository(db *sql.DB) repository.ClubRepositoryInterface {
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

func (r ClubRepository) CreateClub(club *entity.Club) (*entity.Club, errs.AppErrorInterface) {
	result, err := r.db.Exec("INSERT INTO clubs (name, status_id) VALUES (?, ?)", club.Name.Value, club.Status.Id)

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	id, err := result.LastInsertId()

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	newId, err := value.NewId(uint(id))

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	club.Id = *newId

	return club, nil
}

func (r ClubRepository) FindByName(clubName string) (*entity.Club, errs.AppErrorInterface) {
	var id uint
	var name string
	var statusId int
	err := r.db.QueryRow("SELECT id, name, status_id FROM clubs where name = ?", clubName).Scan(&id, &name, &statusId)

	if errors.Is(err, sql.ErrNoRows) {
		return &entity.Club{}, nil
	}

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	club, err := entity.NewClub(id, name, statusId)

	if err != nil {
		return nil, errs.NewAppError(http.StatusNotFound, err.Error())
	}

	return club, nil
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
