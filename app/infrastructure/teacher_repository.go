package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type TeacherRepository struct {
	db *sql.DB
}

// レポジトリの生成
func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r TeacherRepository) GetTeachers() ([]*entity.Teacher, errs.AppErrorInterface) {
	rows, err := r.db.Query("SELECT id, name from teachers")

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	var teachers []*entity.Teacher
	for rows.Next() {
		teacher, err := r.scanTeacher(rows)

		if err != nil {
			return nil, err
		}

		teachers = append(teachers, teacher)
	}

	err = rows.Err()

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return teachers, nil
}

func (r TeacherRepository) scanTeacher(rows *sql.Rows) (*entity.Teacher, errs.AppErrorInterface) {
	var id uint
	var name string
	err := rows.Scan(&id, &name)
	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	teacher, err := entity.NewTeacher(id, name)

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return teacher, nil
}
