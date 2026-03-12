package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/models"
	"golang-entrypoint/internal/repository"
)

type CourseHandler struct {
	repo *repository.CourseRepository
}

func NewCourseHandler(repo *repository.CourseRepository) *CourseHandler {
	return &CourseHandler{repo: repo}
}

func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c models.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if c.Title == "" || c.TeacherID <= 0 {
		writeError(w, http.StatusBadRequest, "missing or invalid required fields")
		return
	}

	if err := h.repo.CreateCourse(&c); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create course")
		return
	}

	writeJSON(w, http.StatusCreated, c)
}

func (h *CourseHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid course ID")
		return
	}

	c, err := h.repo.GetCourseByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "course not found")
		return
	}

	writeJSON(w, http.StatusOK, c)
}

func (h *CourseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	courses, err := h.repo.GetAllCourses()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch courses")
		return
	}

	if courses == nil {
		courses = []models.Course{}
	}

	writeJSON(w, http.StatusOK, courses)
}

func (h *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid course ID")
		return
	}

	var c models.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	c.ID = id

	if err := h.repo.UpdateCourse(&c); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update course")
		return
	}

	writeJSON(w, http.StatusOK, c)
}

func (h *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid course ID")
		return
	}

	if err := h.repo.DeleteCourse(id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete course")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
