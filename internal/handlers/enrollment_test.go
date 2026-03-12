package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEnrollmentService struct{}

func (m *mockEnrollmentService) Enroll(_, _ int) error   { return nil }
func (m *mockEnrollmentService) Unenroll(_, _ int) error { return nil }

func TestEnrollmentHandler_Enroll_InvalidIDs(t *testing.T) {
	h := NewEnrollmentHandler(&mockEnrollmentService{})

	req := httptest.NewRequest(http.MethodPost, "/students/abc/courses/def", nil)
	rec := httptest.NewRecorder()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /students/{id}/courses/{course_id}", h.Enroll)

	mux.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid IDs")
}
