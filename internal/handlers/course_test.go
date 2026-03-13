package handlers

import (
	"bytes"
	"context"
	"golang-entrypoint/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-entrypoint/internal/domain"

	"github.com/stretchr/testify/assert"
)

type mockCourseService struct{}

func (m *mockCourseService) CreateCourse(_ context.Context, _ *domain.Course) (*domain.Course, error) {
	return nil, nil
}
func (m *mockCourseService) GetCourseByID(_ context.Context, _ int) (*domain.Course, error) {
	return nil, nil
}
func (m *mockCourseService) GetAllCourses(_ context.Context) ([]domain.Course, error) {
	return nil, nil
}
func (m *mockCourseService) UpdateCourse(_ context.Context, _ *domain.Course) (*domain.Course, error) {
	return nil, nil
}
func (m *mockCourseService) DeleteCourse(_ context.Context, _ int) (int, error) { return 0, nil }

func TestCourseHandler_Create_InvalidJSON(t *testing.T) {
	h := NewCourseHandler(&mockCourseService{}, &mock.ValidatorService{})

	req := httptest.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer([]byte(`{wrong json format`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}
