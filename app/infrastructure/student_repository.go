package infrastructure

import (
	"database/sql"
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/errs"
)

type StudentRepository struct {
	db *sql.DB
}

// レポジトリの生成
func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r StudentRepository) GetStudents() ([]*entity.Student, errs.AppErrorInterface) {
	rows, err := r.db.Query("SELECT id, name from students")

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	defer rows.Close()

	var students []*entity.Student
	for rows.Next() {
		student, err := r.scanStudent(rows)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	err = rows.Err()

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return students, nil
}

func (r StudentRepository) Find(studentId int) (*entity.Student, errs.AppErrorInterface) {
	var id uint
	var name string
	r.db.QueryRow("SELECT id, name FROM students where id = ?", studentId).Scan(&id, &name)

	student, err := entity.NewStudent(id, name)

	if err != nil {
		return nil, errs.NewAppError(http.StatusNotFound, err.Error())
	}

	return student, nil
}

func (r StudentRepository) scanStudent(rows *sql.Rows) (*entity.Student, errs.AppErrorInterface) {
	var id uint
	var name string
	err := rows.Scan(&id, &name)
	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	student, err := entity.NewStudent(id, name)

	if err != nil {
		return nil, errs.NewAppError(http.StatusInternalServerError, err.Error())
	}

	return student, nil
}
