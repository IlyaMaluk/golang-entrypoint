package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-entrypoint/internal/domain"

	"github.com/stretchr/testify/assert"
)

type mockStudentService struct{}

func (m *mockStudentService) CreateStudent(_ *domain.Student) (*domain.Student, error) {
	return nil, nil
}
func (m *mockStudentService) GetStudentByID(_ int) (*domain.Student, error) { return nil, nil }
func (m *mockStudentService) GetAllStudents() ([]domain.Student, error)     { return nil, nil }
func (m *mockStudentService) UpdateStudent(_ *domain.Student) (*domain.Student, error) {
	return nil, nil
}
func (m *mockStudentService) DeleteStudent(_ int) (int, error) { return 0, nil }

func TestStudentHandler_Create_InvalidJSON(t *testing.T) {
	h := NewStudentHandler(&mockStudentService{})

	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer([]byte(`{bad json`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestStudentHandler_Create_MissingFields(t *testing.T) {
	h := NewStudentHandler(&mockStudentService{})

	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer([]byte(`{"first_name": "John"}`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "missing required fields")
}
