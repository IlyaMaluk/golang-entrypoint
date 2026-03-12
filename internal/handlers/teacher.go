package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/models"
	"golang-entrypoint/internal/repository"
)

type TeacherHandler struct {
	repo *repository.TeacherRepository
}

func NewTeacherHandler(repo *repository.TeacherRepository) *TeacherHandler {
	return &TeacherHandler{repo: repo}
}

func (h *TeacherHandler) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if t.FirstName == "" || t.LastName == "" || t.Department == "" {
		writeError(w, http.StatusBadRequest, "missing required fields")
		return
	}

	if err := h.repo.CreateTeacher(&t); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create teacher")
		return
	}

	writeJSON(w, http.StatusCreated, t)
}

func (h *TeacherHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid teacher ID")
		return
	}

	t, err := h.repo.GetTeacherByID(id)
	if err != nil {
		writeError(w, http.StatusNotFound, "teacher not found")
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *TeacherHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	teachers, err := h.repo.GetAllTeachers()
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to fetch teachers")
		return
	}

	if teachers == nil {
		teachers = []models.Teacher{}
	}

	writeJSON(w, http.StatusOK, teachers)
}

func (h *TeacherHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid teacher ID")
		return
	}

	var t models.Teacher
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	t.ID = id

	if err := h.repo.UpdateTeacher(&t); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to update teacher")
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func (h *TeacherHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid teacher ID")
		return
	}

	if err := h.repo.DeleteTeacher(id); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to delete teacher")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
