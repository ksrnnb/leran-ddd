package factory

import (
	"net/http"

	"github.com/ksrnnb/learn-ddd/domain/entity"
	"github.com/ksrnnb/learn-ddd/domain/repository"
	"github.com/ksrnnb/learn-ddd/domain/value"
	"github.com/ksrnnb/learn-ddd/errs"
)

type ClubFactory struct {
	clubRepo    repository.ClubRepositoryInterface
	studnetRepo repository.StudentRepositoryInterface
	teacherRepo repository.TeacherRepositoryInterface
}

func NewClubFactory(repo repository.ClubRepositoryInterface) *ClubFactory {
	return &ClubFactory{clubRepo: repo}
}

func (f ClubFactory) NewClubWithStudentsAndTeacher(name string, studentIds []int, teacherId int) (*entity.Club, errs.AppErrorInterface) {
	nameObject, err := value.NewClubName(name)

	if err != nil {
		return nil, errs.NewAppError(http.StatusUnprocessableEntity, err.Error())
	}

	status := f.newStatus(studentIds)

	students, appErr := f.newStudents(studentIds)

	if appErr != nil {
		return nil, appErr
	}

	teacher, appErr := f.newTeacher(teacherId)

	if appErr != nil {
		return nil, appErr
	}

	return &entity.Club{
		Name:     *nameObject,
		Status:   *status,
		Students: students,
		Teacher:  teacher,
	}, nil
}

func (f ClubFactory) newStatus(studentIds []int) *value.ClubStatus {
	if len(studentIds) < entity.StudentsLowerLimit {
		return value.NewUnapprovedStatus()
	}

	return value.NewApprovedStatus()
}

func (f ClubFactory) newStudents(studentIds []int) ([]*entity.Student, errs.AppErrorInterface) {
	var students []*entity.Student

	// getStudentsでとってきて、そこでループ回してリクエストを減らす。
	for _, studentId := range studentIds {
		student, err := f.studnetRepo.Find(studentId)

		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	return students, nil
}

func (f ClubFactory) newTeacher(teacherId int) (*entity.Teacher, errs.AppErrorInterface) {
	return f.teacherRepo.Find(teacherId)
}
