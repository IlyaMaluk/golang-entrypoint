package repository

import (
	"context"
	"database/sql"
	"golang-entrypoint/internal/models"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(ctx context.Context, s *models.Student) (*models.Student, error) {
	query := `INSERT INTO students (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id, first_name, last_name, email`
	created := &models.Student{}
	err := r.db.QueryRowContext(ctx, query, s.FirstName, s.LastName, s.Email).Scan(
		&created.ID, &created.FirstName, &created.LastName, &created.Email,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *StudentRepository) GetStudentByID(ctx context.Context, id int) (*models.Student, error) {
	query := `SELECT id, first_name, last_name, email FROM students WHERE id = $1`
	s := &models.Student{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&s.ID, &s.FirstName, &s.LastName, &s.Email)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StudentRepository) GetAllStudents(ctx context.Context) ([]models.Student, error) {
	query := `SELECT id, first_name, last_name, email FROM students`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var s models.Student
		if err := rows.Scan(&s.ID, &s.FirstName, &s.LastName, &s.Email); err != nil {
			return nil, err
		}
		students = append(students, s)
	}
	return students, nil
}

func (r *StudentRepository) UpdateStudent(ctx context.Context, s *models.Student) (*models.Student, error) {
	query := `UPDATE students SET first_name = $1, last_name = $2, email = $3 WHERE id = $4 RETURNING id, first_name, last_name, email`
	updated := &models.Student{}
	err := r.db.QueryRowContext(ctx, query, s.FirstName, s.LastName, s.Email, s.ID).Scan(
		&updated.ID, &updated.FirstName, &updated.LastName, &updated.Email,
	)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *StudentRepository) DeleteStudent(ctx context.Context, id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
