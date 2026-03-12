package repository

import (
	"database/sql"
	"golang-entrypoint/internal/models"
)

type TeacherRepository struct {
	db *sql.DB
}

func NewTeacherRepository(db *sql.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) CreateTeacher(t *models.Teacher) error {
	query := `INSERT INTO teachers (first_name, last_name, department) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(query, t.FirstName, t.LastName, t.Department).Scan(&t.ID)
}

func (r *TeacherRepository) GetTeacherByID(id int) (*models.Teacher, error) {
	query := `SELECT id, first_name, last_name, department FROM teachers WHERE id = $1`
	t := &models.Teacher{}
	err := r.db.QueryRow(query, id).Scan(&t.ID, &t.FirstName, &t.LastName, &t.Department)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (r *TeacherRepository) GetAllTeachers() ([]models.Teacher, error) {
	query := `SELECT id, first_name, last_name, department FROM teachers`
	rows, err := r.db.Query(query)
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

func (r *TeacherRepository) UpdateTeacher(t *models.Teacher) error {
	query := `UPDATE teachers SET first_name = $1, last_name = $2, department = $3 WHERE id = $4`
	_, err := r.db.Exec(query, t.FirstName, t.LastName, t.Department, t.ID)
	return err
}

func (r *TeacherRepository) DeleteTeacher(id int) error {
	query := `DELETE FROM teachers WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
