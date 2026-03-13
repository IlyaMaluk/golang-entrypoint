package service

import (
	"context"
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type TeacherRepository interface {
	CreateTeacher(ctx context.Context, t *models.Teacher) (*models.Teacher, error)
	GetTeacherByID(ctx context.Context, id int) (*models.Teacher, error)
	GetAllTeachers(ctx context.Context) ([]models.Teacher, error)
	UpdateTeacher(ctx context.Context, t *models.Teacher) (*models.Teacher, error)
	DeleteTeacher(ctx context.Context, id int) error
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

func (s *TeacherServiceImpl) CreateTeacher(ctx context.Context, req *domain.Teacher) (*domain.Teacher, error) {
	dbModel := mapTeacherToDB(req)
	createdDBModel, err := s.repo.CreateTeacher(ctx, dbModel)
	if err != nil {
		return nil, err
	}
	return mapTeacherToDomain(createdDBModel), nil
}

func (s *TeacherServiceImpl) GetTeacherByID(ctx context.Context, id int) (*domain.Teacher, error) {
	dbModel, err := s.repo.GetTeacherByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapTeacherToDomain(dbModel), nil
}

func (s *TeacherServiceImpl) GetAllTeachers(ctx context.Context) ([]domain.Teacher, error) {
	dbModels, err := s.repo.GetAllTeachers(ctx)
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

func (s *TeacherServiceImpl) UpdateTeacher(ctx context.Context, req *domain.Teacher) (*domain.Teacher, error) {
	updatedDBModel, err := s.repo.UpdateTeacher(ctx, mapTeacherToDB(req))
	if err != nil {
		return nil, err
	}
	return mapTeacherToDomain(updatedDBModel), nil
}

func (s *TeacherServiceImpl) DeleteTeacher(ctx context.Context, id int) (int, error) {
	if err := s.repo.DeleteTeacher(ctx, id); err != nil {
		return 0, err
	}
	return id, nil
}
