package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/models"
	"golang-entrypoint/internal/repository"
)

type StudentHandler struct {
	repo *repository.StudentRepository
}

func NewStudentHandler(repo *repository.StudentRepository) *StudentHandler {
	return &StudentHandler{repo: repo}
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if s.FirstName == "" || s.LastName == "" || s.Email == "" {
		writeError(w, http.StatusBadRequest, "missing required fields")
		return
	}

	if err := h.repo.CreateStudent(&s); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create student")
		return
	}

	writeJSON(w, http.StatusCreated, s)
}

func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid student ID")
		return
	}

	s, err := h.repo.GetStudentByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "student not found")
		return
	}

	writeJSON(w, http.StatusOK, s)
}

func (h *StudentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	students, err := h.repo.GetAllStudents()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch students")
		return
	}

	if students == nil {
		students = []models.Student{}
	}

	writeJSON(w, http.StatusOK, students)
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid student ID")
		return
	}

	var s models.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	s.ID = id

	if err := h.repo.UpdateStudent(&s); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update student")
		return
	}

	writeJSON(w, http.StatusOK, s)
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid student ID")
		return
	}

	if err := h.repo.DeleteStudent(id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete student")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
