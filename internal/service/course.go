package service

import (
	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/models"
	"golang-entrypoint/internal/repository"
)

type CourseService struct {
	repo *repository.CourseRepository
}

func NewCourseService(repo *repository.CourseRepository) *CourseService {
	return &CourseService{repo: repo}
}

// Функції для конвертації моделей
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

// Бізнес-логіка
func (s *CourseService) CreateCourse(req *domain.Course) error {
	dbModel := mapCourseToDB(req)

	if err := s.repo.CreateCourse(dbModel); err != nil {
		return err
	}

	req.ID = dbModel.ID
	return nil
}

func (s *CourseService) GetCourseByID(id int) (*domain.Course, error) {
	dbModel, err := s.repo.GetCourseByID(id)
	if err != nil {
		return nil, err
	}

	return mapCourseToDomain(dbModel), nil
}

func (s *CourseService) GetAllCourses() ([]domain.Course, error) {
	dbModels, err := s.repo.GetAllCourses()
	if err != nil {
		return nil, err
	}

	var domainCourses []domain.Course
	for _, dbModel := range dbModels {
		domainCourses = append(domainCourses, *mapCourseToDomain(&dbModel))
	}

	if domainCourses == nil {
		domainCourses = []domain.Course{} // Щоб в JSON повертався [] замість null
	}

	return domainCourses, nil
}

func (s *CourseService) UpdateCourse(req *domain.Course) error {
	return s.repo.UpdateCourse(mapCourseToDB(req))
}

func (s *CourseService) DeleteCourse(id int) error {
	return s.repo.DeleteCourse(id)
}
