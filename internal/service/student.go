package service

import (
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type StudentRepository interface {
	CreateStudent(s *models.Student) (int, error)
	GetStudentByID(id int) (*models.Student, error)
	GetAllStudents() ([]models.Student, error)
	UpdateStudent(s *models.Student) error
	DeleteStudent(id int) error
}

type StudentServiceImpl struct {
	repo StudentRepository
}

func NewStudentService(repo StudentRepository) *StudentServiceImpl {
	return &StudentServiceImpl{repo: repo}
}

func mapStudentToDB(d *domain.Student) *models.Student {
	return &models.Student{
		ID:        d.ID,
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Email:     d.Email,
	}
}

func mapStudentToDomain(m *models.Student) *domain.Student {
	return &domain.Student{
		ID:        m.ID,
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
	}
}

func (s *StudentServiceImpl) CreateStudent(req *domain.Student) (*domain.Student, error) {
	dbModel := mapStudentToDB(req)
	id, err := s.repo.CreateStudent(dbModel)
	if err != nil {
		return nil, err
	}
	res := *req
	res.ID = id
	return &res, nil
}

func (s *StudentServiceImpl) GetStudentByID(id int) (*domain.Student, error) {
	dbModel, err := s.repo.GetStudentByID(id)
	if err != nil {
		return nil, err
	}
	return mapStudentToDomain(dbModel), nil
}

func (s *StudentServiceImpl) GetAllStudents() ([]domain.Student, error) {
	dbModels, err := s.repo.GetAllStudents()
	if err != nil {
		return nil, err
	}

	var domainStudents []domain.Student
	for _, dbModel := range dbModels {
		domainStudents = append(domainStudents, *mapStudentToDomain(&dbModel))
	}
	if domainStudents == nil {
		domainStudents = []domain.Student{}
	}
	return domainStudents, nil
}

func (s *StudentServiceImpl) UpdateStudent(req *domain.Student) (*domain.Student, error) {
	if err := s.repo.UpdateStudent(mapStudentToDB(req)); err != nil {
		return nil, err
	}
	res := *req
	return &res, nil
}

func (s *StudentServiceImpl) DeleteStudent(id int) (int, error) {
	if err := s.repo.DeleteStudent(id); err != nil {
		return 0, err
	}
	return id, nil
}
