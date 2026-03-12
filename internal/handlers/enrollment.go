package handlers

import (
	"net/http"
	"strconv"

	"golang-entrypoint/internal/repository"
)

type EnrollmentHandler struct {
	repo *repository.EnrollmentRepository
}

func NewEnrollmentHandler(repo *repository.EnrollmentRepository) *EnrollmentHandler {
	return &EnrollmentHandler{repo: repo}
}

func (h *EnrollmentHandler) Enroll(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.PathValue("id")
	courseIDStr := r.PathValue("course_id")

	studentID, err1 := strconv.Atoi(studentIDStr)
	courseID, err2 := strconv.Atoi(courseIDStr)

	if err1 != nil || err2 != nil {
		writeError(w, http.StatusBadRequest, "invalid IDs")
		return
	}

	if err := h.repo.EnrollStudent(studentID, courseID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to enroll student")
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *EnrollmentHandler) Unenroll(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.PathValue("id")
	courseIDStr := r.PathValue("course_id")

	studentID, err1 := strconv.Atoi(studentIDStr)
	courseID, err2 := strconv.Atoi(courseIDStr)

	if err1 != nil || err2 != nil {
		writeError(w, http.StatusBadRequest, "invalid IDs")
		return
	}

	if err := h.repo.UnenrollStudent(studentID, courseID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to unenroll student")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
