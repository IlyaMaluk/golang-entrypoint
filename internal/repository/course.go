package repository

import (
	"context"
	"database/sql"
	"golang-entrypoint/internal/models"
)

type CourseRepository struct {
	db *sql.DB
}

func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (r *CourseRepository) CreateCourse(ctx context.Context, c *models.Course) (*models.Course, error) {
	query := `INSERT INTO courses (title, description, teacher_id) VALUES ($1, $2, $3) RETURNING id, title, description, teacher_id`
	created := &models.Course{}
	err := r.db.QueryRowContext(ctx, query, c.Title, c.Description, c.TeacherID).Scan(
		&created.ID, &created.Title, &created.Description, &created.TeacherID,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *CourseRepository) GetCourseByID(ctx context.Context, id int) (*models.Course, error) {
	query := `SELECT id, title, description, teacher_id FROM courses WHERE id = $1`
	c := &models.Course{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&c.ID, &c.Title, &c.Description, &c.TeacherID)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *CourseRepository) GetAllCourses(ctx context.Context) ([]models.Course, error) {
	query := `SELECT id, title, description, teacher_id FROM courses`
	rows, err := r.db.QueryContext(ctx, query)
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

func (r *CourseRepository) UpdateCourse(ctx context.Context, c *models.Course) (*models.Course, error) {
	query := `UPDATE courses SET title = $1, description = $2, teacher_id = $3 WHERE id = $4 RETURNING id, title, description, teacher_id`
	updated := &models.Course{}
	err := r.db.QueryRowContext(ctx, query, c.Title, c.Description, c.TeacherID, c.ID).Scan(
		&updated.ID, &updated.Title, &updated.Description, &updated.TeacherID,
	)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *CourseRepository) DeleteCourse(ctx context.Context, id int) error {
	query := `DELETE FROM courses WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
