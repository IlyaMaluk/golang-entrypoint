package repository

import (
	"database/sql"
	"golang-entrypoint/internal/models"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{db: db}
}

func (r *StudentRepository) CreateStudent(s *models.Student) (int, error) {
	var id int
	query := `INSERT INTO students (first_name, last_name, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, s.FirstName, s.LastName, s.Email).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *StudentRepository) GetStudentByID(id int) (*models.Student, error) {
	query := `SELECT id, first_name, last_name, email FROM students WHERE id = $1`
	s := &models.Student{}
	err := r.db.QueryRow(query, id).Scan(&s.ID, &s.FirstName, &s.LastName, &s.Email)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StudentRepository) GetAllStudents() ([]models.Student, error) {
	query := `SELECT id, first_name, last_name, email FROM students`
	rows, err := r.db.Query(query)
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

func (r *StudentRepository) UpdateStudent(s *models.Student) error {
	query := `UPDATE students SET first_name = $1, last_name = $2, email = $3 WHERE id = $4`
	_, err := r.db.Exec(query, s.FirstName, s.LastName, s.Email, s.ID)
	return err
}

func (r *StudentRepository) DeleteStudent(id int) error {
	query := `DELETE FROM students WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}
