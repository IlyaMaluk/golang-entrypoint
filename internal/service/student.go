package service

import (
	"context"
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type StudentRepository interface {
	CreateStudent(ctx context.Context, s *models.Student) (*models.Student, error)
	GetStudentByID(ctx context.Context, id int) (*models.Student, error)
	GetAllStudents(ctx context.Context) ([]models.Student, error)
	UpdateStudent(ctx context.Context, s *models.Student) (*models.Student, error)
	DeleteStudent(ctx context.Context, id int) error
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

func (s *StudentServiceImpl) CreateStudent(ctx context.Context, req *domain.Student) (*domain.Student, error) {
	dbModel := mapStudentToDB(req)
	createdDBModel, err := s.repo.CreateStudent(ctx, dbModel)
	if err != nil {
		return nil, err
	}
	return mapStudentToDomain(createdDBModel), nil
}

func (s *StudentServiceImpl) GetStudentByID(ctx context.Context, id int) (*domain.Student, error) {
	dbModel, err := s.repo.GetStudentByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapStudentToDomain(dbModel), nil
}

func (s *StudentServiceImpl) GetAllStudents(ctx context.Context) ([]domain.Student, error) {
	dbModels, err := s.repo.GetAllStudents(ctx)
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

func (s *StudentServiceImpl) UpdateStudent(ctx context.Context, req *domain.Student) (*domain.Student, error) {
	updatedDBModel, err := s.repo.UpdateStudent(ctx, mapStudentToDB(req))
	if err != nil {
		return nil, err
	}
	return mapStudentToDomain(updatedDBModel), nil
}

func (s *StudentServiceImpl) DeleteStudent(ctx context.Context, id int) (int, error) {
	if err := s.repo.DeleteStudent(ctx, id); err != nil {
		return 0, err
	}
	return id, nil
}
