package handlers

import (
	"net/http"
	"strconv"
)

type EnrollmentService interface {
	Enroll(studentID, courseID int) error
	Unenroll(studentID, courseID int) error
}

type EnrollmentHandler struct {
	svc EnrollmentService
}

func NewEnrollmentHandler(svc EnrollmentService) *EnrollmentHandler {
	return &EnrollmentHandler{svc: svc}
}

func (h *EnrollmentHandler) Enroll(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.PathValue("id")
	courseIDStr := r.PathValue("course_id")

	studentID, err1 := strconv.Atoi(studentIDStr)
	courseID, err2 := strconv.Atoi(courseIDStr)

	if err1 != nil || err2 != nil {
		writeJSONError(w, "invalid IDs", http.StatusBadRequest)
		return
	}

	if err := h.svc.Enroll(studentID, courseID); err != nil {
		writeJSONError(w, "failed to enroll student", http.StatusInternalServerError)
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
		writeJSONError(w, "invalid IDs", http.StatusBadRequest)
		return
	}

	if err := h.svc.Unenroll(studentID, courseID); err != nil {
		writeJSONError(w, "failed to unenroll student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
