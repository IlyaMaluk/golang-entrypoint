package service

import (
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type TeacherRepository interface {
	CreateTeacher(t *models.Teacher) (int, error)
	GetTeacherByID(id int) (*models.Teacher, error)
	GetAllTeachers() ([]models.Teacher, error)
	UpdateTeacher(t *models.Teacher) (*models.Teacher, error)
	DeleteTeacher(id int) error
}

type TeacherServiceImpl struct {
	repo TeacherRepository
}

func NewTeacherService(repo TeacherRepository) *TeacherServiceImpl {
	return &TeacherServiceImpl{repo: repo}
}

func mapTeacherToDB(d *domain.Teacher) *models.Teacher {
	return &models.Teacher{
		ID:         d.ID,
		FirstName:  d.FirstName,
		LastName:   d.LastName,
		Department: d.Department,
	}
}

func mapTeacherToDomain(m *models.Teacher) *domain.Teacher {
	return &domain.Teacher{
		ID:         m.ID,
		FirstName:  m.FirstName,
		LastName:   m.LastName,
		Department: m.Department,
	}
}

func (s *TeacherServiceImpl) CreateTeacher(req *domain.Teacher) (*domain.Teacher, error) {
	dbModel := mapTeacherToDB(req)
	id, err := s.repo.CreateTeacher(dbModel)
	if err != nil {
		return nil, err
	}

	return &domain.Teacher{
		ID:         id,
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Department: req.Department,
	}, nil
}

func (s *TeacherServiceImpl) GetTeacherByID(id int) (*domain.Teacher, error) {
	dbModel, err := s.repo.GetTeacherByID(id)
	if err != nil {
		return nil, err
	}
	return mapTeacherToDomain(dbModel), nil
}

func (s *TeacherServiceImpl) GetAllTeachers() ([]domain.Teacher, error) {
	dbModels, err := s.repo.GetAllTeachers()
	if err != nil {
		return nil, err
	}

	var domainTeachers []domain.Teacher
	for _, dbModel := range dbModels {
		domainTeachers = append(domainTeachers, *mapTeacherToDomain(&dbModel))
	}
	if domainTeachers == nil {
		domainTeachers = []domain.Teacher{}
	}
	return domainTeachers, nil
}

func (s *TeacherServiceImpl) UpdateTeacher(req *domain.Teacher) (*domain.Teacher, error) {
	updatedDBModel, err := s.repo.UpdateTeacher(mapTeacherToDB(req))
	if err != nil {
		return nil, err
	}
	return mapTeacherToDomain(updatedDBModel), nil
}

func (s *TeacherServiceImpl) DeleteTeacher(id int) (int, error) {
	if err := s.repo.DeleteTeacher(id); err != nil {
		return 0, err
	}
	return id, nil
}
