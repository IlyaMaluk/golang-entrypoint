package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-entrypoint/internal/repository"

	"github.com/stretchr/testify/assert"
)

func TestStudentHandler_Create_InvalidJSON(t *testing.T) {
	h := NewStudentHandler(&repository.StudentRepository{})

	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer([]byte(`{bad json`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "invalid request body")
}

func TestStudentHandler_Create_MissingFields(t *testing.T) {
	h := NewStudentHandler(&repository.StudentRepository{})

	req := httptest.NewRequest(http.MethodPost, "/students", bytes.NewBuffer([]byte(`{"first_name": "John"}`)))
	rec := httptest.NewRecorder()

	h.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "missing required fields")
}
