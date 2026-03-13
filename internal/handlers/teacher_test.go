package handlers

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-entrypoint/internal/domain"

	"github.com/stretchr/testify/assert"
)

type mockTeacherService struct{}

func (m *mockTeacherService) CreateTeacher(_ context.Context, _ *domain.Teacher) (*domain.Teacher, error) {
	return nil, nil
}
func (m *mockTeacherService) GetTeacherByID(_ context.Context, _ int) (*domain.Teacher, error) {
	return nil, nil
}
func (m *mockTeacherService) GetAllTeachers(_ context.Context) ([]domain.Teacher, error) {
	return nil, nil
}
func (m *mockTeacherService) UpdateTeacher(_ context.Context, _ *domain.Teacher) (*domain.Teacher, error) {
	return nil, nil
}
func (m *mockTeacherService) DeleteTeacher(_ context.Context, _ int) (int, error) { return 0, nil }

func TestTeacherHandler_Create_MissingFields(t *testing.T) {
	h := NewTeacherHandler(&mockTeacherService{})

	req := httptest.NewRequest(http.MethodPost, "/teachers", bytes.NewBuffer([]byte(`{"first_name": "Linus"}`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "missing required fields")
}
