package service

import (
	"context"
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type CourseRepository interface {
	CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error)
	GetCourseByID(ctx context.Context, id int) (*models.Course, error)
	GetAllCourses(ctx context.Context) ([]models.Course, error)
	UpdateCourse(ctx context.Context, c *models.Course) (*models.Course, error)
	DeleteCourse(ctx context.Context, id int) error
}

type CourseServiceImpl struct {
	repo CourseRepository
}

func NewCourseService(repo CourseRepository) *CourseServiceImpl {
	return &CourseServiceImpl{repo: repo}
}

func mapCourseToDB(d *domain.Course) *models.Course {
	return &models.Course{
		ID:          d.ID,
		Title:       d.Title,
		Description: d.Description,
		TeacherID:   d.TeacherID,
	}
}

func mapCourseToDomain(m *models.Course) *domain.Course {
	return &domain.Course{
		ID:          m.ID,
		Title:       m.Title,
		Description: m.Description,
		TeacherID:   m.TeacherID,
	}
}

func (s *CourseServiceImpl) CreateCourse(ctx context.Context, req *domain.Course) (*domain.Course, error) {
	dbModel := mapCourseToDB(req)
	createdDBModel, err := s.repo.CreateCourse(ctx, dbModel)
	if err != nil {
		return nil, err
	}
	return mapCourseToDomain(createdDBModel), nil
}

func (s *CourseServiceImpl) GetCourseByID(ctx context.Context, id int) (*domain.Course, error) {
	dbModel, err := s.repo.GetCourseByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapCourseToDomain(dbModel), nil
}

func (s *CourseServiceImpl) GetAllCourses(ctx context.Context) ([]domain.Course, error) {
	dbModels, err := s.repo.GetAllCourses(ctx)
	if err != nil {
		return nil, err
	}

	var domainCourses []domain.Course
	for _, dbModel := range dbModels {
		domainCourses = append(domainCourses, *mapCourseToDomain(&dbModel))
	}
	if domainCourses == nil {
		domainCourses = []domain.Course{}
	}
	return domainCourses, nil
}

func (s *CourseServiceImpl) UpdateCourse(ctx context.Context, req *domain.Course) (*domain.Course, error) {
	updatedDBModel, err := s.repo.UpdateCourse(ctx, mapCourseToDB(req))
	if err != nil {
		return nil, err
	}
	return mapCourseToDomain(updatedDBModel), nil
}

func (s *CourseServiceImpl) DeleteCourse(ctx context.Context, id int) (int, error) {
	if err := s.repo.DeleteCourse(ctx, id); err != nil {
		return 0, err
	}
	return id, nil
}
