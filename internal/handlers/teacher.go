package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
)

type TeacherService interface {
	CreateTeacher(req *domain.Teacher) (*domain.Teacher, error)
	GetTeacherByID(id int) (*domain.Teacher, error)
	GetAllTeachers() ([]domain.Teacher, error)
	UpdateTeacher(req *domain.Teacher) (*domain.Teacher, error)
	DeleteTeacher(id int) (int, error)
}

type TeacherHandler struct {
	svc TeacherService
}

func NewTeacherHandler(svc TeacherService) *TeacherHandler {
	return &TeacherHandler{svc: svc}
}

func (h *TeacherHandler) Create(w http.ResponseWriter, r *http.Request) {
	var t domain.Teacher
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if t.FirstName == "" || t.LastName == "" || t.Department == "" {
		writeJSONError(w, "missing required fields", http.StatusBadRequest)
		return
	}

	createdTeacher, err := h.svc.CreateTeacher(&t)
	if err != nil {
		writeJSONError(w, "failed to create teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdTeacher); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *TeacherHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid teacher ID", http.StatusBadRequest)
		return
	}

	t, err := h.svc.GetTeacherByID(id)
	if err != nil {
		writeJSONError(w, "teacher not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *TeacherHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	teachers, err := h.svc.GetAllTeachers()
	if err != nil {
		writeJSONError(w, "failed to fetch teachers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(teachers); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *TeacherHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid teacher ID", http.StatusBadRequest)
		return
	}

	var t domain.Teacher
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	t.ID = id

	updatedTeacher, err := h.svc.UpdateTeacher(&t)
	if err != nil {
		writeJSONError(w, "failed to update teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedTeacher); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *TeacherHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid teacher ID", http.StatusBadRequest)
		return
	}

	deletedID, err := h.svc.DeleteTeacher(id)
	if err != nil {
		writeJSONError(w, "failed to delete teacher", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]int{"deleted_id": deletedID}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}
