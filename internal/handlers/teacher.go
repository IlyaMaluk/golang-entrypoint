package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
)

type TeacherService interface {
	CreateTeacher(ctx context.Context, req *domain.Teacher) (*domain.Teacher, error)
	GetTeacherByID(ctx context.Context, id int) (*domain.Teacher, error)
	GetAllTeachers(ctx context.Context) ([]domain.Teacher, error)
	UpdateTeacher(ctx context.Context, req *domain.Teacher) (*domain.Teacher, error)
	DeleteTeacher(ctx context.Context, id int) (int, error)
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

	ctx := r.Context()
	createdTeacher, err := h.svc.CreateTeacher(ctx, &t)
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

	ctx := r.Context()
	t, err := h.svc.GetTeacherByID(ctx, id)
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

func (h *TeacherHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	teachers, err := h.svc.GetAllTeachers(ctx)
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

	ctx := r.Context()
	updatedTeacher, err := h.svc.UpdateTeacher(ctx, &t)
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

	ctx := r.Context()
	deletedID, err := h.svc.DeleteTeacher(ctx, id)
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
