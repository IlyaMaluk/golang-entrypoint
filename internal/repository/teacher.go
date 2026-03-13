package repository

import (
	"context"
	"database/sql"
	"golang-entrypoint/internal/models"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) CreateTeacher(ctx context.Context, t *models.Teacher) (*models.Teacher, error) {
	query := `INSERT INTO teachers (first_name, last_name, department) VALUES ($1, $2, $3) RETURNING id, first_name, last_name, department`
	created := &models.Teacher{}
	err := r.db.QueryRowContext(ctx, query, t.FirstName, t.LastName, t.Department).Scan(
		&created.ID, &created.FirstName, &created.LastName, &created.Department,
	)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (r *TeacherRepository) GetTeacherByID(ctx context.Context, id int) (*models.Teacher, error) {
	query := `SELECT id, first_name, last_name, department FROM teachers WHERE id = $1`
	t := &models.Teacher{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.FirstName, &t.LastName, &t.Department)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TeacherRepository) GetAllTeachers(ctx context.Context) ([]models.Teacher, error) {
	query := `SELECT id, first_name, last_name, department FROM teachers`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var t models.Teacher
		if err := rows.Scan(&t.ID, &t.FirstName, &t.LastName, &t.Department); err != nil {
			return nil, err
		}
		teachers = append(teachers, t)
	}
	return teachers, nil
}

func (r *TeacherRepository) UpdateTeacher(ctx context.Context, t *models.Teacher) (*models.Teacher, error) {
	query := `UPDATE teachers SET first_name = $1, last_name = $2, department = $3 WHERE id = $4 RETURNING id, first_name, last_name, department`
	updated := &models.Teacher{}
	err := r.db.QueryRowContext(ctx, query, t.FirstName, t.LastName, t.Department, t.ID).Scan(
		&updated.ID, &updated.FirstName, &updated.LastName, &updated.Department,
	)
	if err != nil {
		return nil, err
	}
	return updated, nil
}

func (r *TeacherRepository) DeleteTeacher(ctx context.Context, id int) error {
	query := `DELETE FROM teachers WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
