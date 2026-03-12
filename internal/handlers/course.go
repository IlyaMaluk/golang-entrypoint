package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
)

type CourseService interface {
	CreateCourse(req *domain.Course) (*domain.Course, error)
	GetCourseByID(id int) (*domain.Course, error)
	GetAllCourses() ([]domain.Course, error)
	UpdateCourse(req *domain.Course) (*domain.Course, error)
	DeleteCourse(id int) (int, error)
}

type CourseHandler struct {
	svc CourseService
}

func NewCourseHandler(svc CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": message}); err != nil {
		slog.Error("failed to write json error response", "error", err)
	}
}

func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c domain.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	createdCourse, err := h.svc.CreateCourse(&c)
	if err != nil {
		writeJSONError(w, "failed to create course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdCourse); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *CourseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid course ID", http.StatusBadRequest)
		return
	}

	c, err := h.svc.GetCourseByID(id)
	if err != nil {
		writeJSONError(w, "course not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *CourseHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	courses, err := h.svc.GetAllCourses()
	if err != nil {
		writeJSONError(w, "failed to fetch courses", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(courses); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid course ID", http.StatusBadRequest)
		return
	}

	var c domain.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	c.ID = id

	updatedCourse, err := h.svc.UpdateCourse(&c)
	if err != nil {
		writeJSONError(w, "failed to update course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedCourse); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid course ID", http.StatusBadRequest)
		return
	}

	deletedID, err := h.svc.DeleteCourse(id)
	if err != nil {
		writeJSONError(w, "failed to delete course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]int{"deleted_id": deletedID}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}
