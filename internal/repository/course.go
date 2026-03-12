package repository

import (
	"database/sql"
	"golang-entrypoint/internal/models"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) CreateCourse(c *models.Course) (int, error) {
	var id int
	query := `INSERT INTO courses (title, description, teacher_id) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, c.Title, c.Description, c.TeacherID).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *CourseRepository) GetCourseByID(id int) (*models.Course, error) {
	query := `SELECT id, title, description, teacher_id FROM courses WHERE id = $1`
	c := &models.Course{}
	err := r.db.QueryRow(query, id).Scan(&c.ID, &c.Title, &c.Description, &c.TeacherID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CourseRepository) GetAllCourses() ([]models.Course, error) {
	query := `SELECT id, title, description, teacher_id FROM courses`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []models.Course
	for rows.Next() {
		var c models.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.TeacherID); err != nil {
			return nil, err
		}
		courses = append(courses, c)
	}
	return courses, nil
}

func (r *CourseRepository) UpdateCourse(c *models.Course) error {
	query := `UPDATE courses SET title = $1, description = $2, teacher_id = $3 WHERE id = $4`
	_, err := r.db.Exec(query, c.Title, c.Description, c.TeacherID, c.ID)
	return err
}

func (r *CourseRepository) DeleteCourse(id int) error {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
