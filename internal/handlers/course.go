package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
	"golang-entrypoint/internal/service"
)

type CourseHandler struct {
	svc *service.CourseService
}

func NewCourseHandler(svc *service.CourseService) *CourseHandler {
	return &CourseHandler{svc: svc}
}

func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c domain.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}

	if err := h.svc.CreateCourse(&c); err != nil {
		http.Error(w, `{"error": "failed to create course"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(c)
	if err != nil {
		return
	}
}

func (h *CourseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid course ID"}`, http.StatusBadRequest)
		return
	}

	c, err := h.svc.GetCourseByID(id)
	if err != nil {
		http.Error(w, `{"error": "course not found"}`, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		return
	}
}

func (h *CourseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	courses, err := h.svc.GetAllCourses()
	if err != nil {
		http.Error(w, `{"error": "failed to fetch courses"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(courses)
	if err != nil {
		return
	}
}

func (h *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid course ID"}`, http.StatusBadRequest)
		return
	}

	var c domain.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, `{"error": "invalid request body"}`, http.StatusBadRequest)
		return
	}
	c.ID = id

	if err := h.svc.UpdateCourse(&c); err != nil {
		http.Error(w, `{"error": "failed to update course"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(c)
	if err != nil {
		return
	}
}

func (h *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error": "invalid course ID"}`, http.StatusBadRequest)
		return
	}

	if err := h.svc.DeleteCourse(id); err != nil {
		http.Error(w, `{"error": "failed to delete course"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
