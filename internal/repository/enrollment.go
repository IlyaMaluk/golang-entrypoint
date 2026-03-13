package repository

import (
	"context"
	"database/sql"
)

type EnrollmentRepository struct {
	db *sql.DB
}

func NewEnrollmentRepository(db *sql.DB) *EnrollmentRepository {
	return &EnrollmentRepository{db: db}
}

func (r *EnrollmentRepository) EnrollStudent(ctx context.Context, studentID, courseID int) error {
	query := `INSERT INTO enrollments (student_id, course_id) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, studentID, courseID)
	return err
}

func (r *EnrollmentRepository) UnenrollStudent(ctx context.Context, studentID, courseID int) error {
	query := `DELETE FROM enrollments WHERE student_id = $1 AND course_id = $2`
	_, err := r.db.ExecContext(ctx, query, studentID, courseID)
	return err
}
