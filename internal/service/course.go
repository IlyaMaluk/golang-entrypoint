package service

import (
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
)

type CourseRepository interface {
	CreateCourse(c *models.Course) (int, error)
	GetCourseByID(id int) (*models.Course, error)
	GetAllCourses() ([]models.Course, error)
	UpdateCourse(c *models.Course) (*models.Course, error)
	DeleteCourse(id int) error
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

func (s *CourseServiceImpl) CreateCourse(req *domain.Course) (*domain.Course, error) {
	dbModel := mapCourseToDB(req)

	id, err := s.repo.CreateCourse(dbModel)
	if err != nil {
		return nil, err
	}

	return &domain.Course{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		TeacherID:   req.TeacherID,
	}, nil
}

func (s *CourseServiceImpl) GetCourseByID(id int) (*domain.Course, error) {
	dbModel, err := s.repo.GetCourseByID(id)
	if err != nil {
		return nil, err
	}
	return mapCourseToDomain(dbModel), nil
}

func (s *CourseServiceImpl) GetAllCourses() ([]domain.Course, error) {
	dbModels, err := s.repo.GetAllCourses()
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

func (s *CourseServiceImpl) UpdateCourse(req *domain.Course) (*domain.Course, error) {
	updatedDBModel, err := s.repo.UpdateCourse(mapCourseToDB(req))
	if err != nil {
		return nil, err
	}
	return mapCourseToDomain(updatedDBModel), nil
}

func (s *CourseServiceImpl) DeleteCourse(id int) (int, error) {
	if err := s.repo.DeleteCourse(id); err != nil {
		return 0, err
	}
	return id, nil
}
