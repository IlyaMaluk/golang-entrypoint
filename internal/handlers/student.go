package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
)

type StudentService interface {
	CreateStudent(req *domain.Student) (*domain.Student, error)
	GetStudentByID(id int) (*domain.Student, error)
	GetAllStudents() ([]domain.Student, error)
	UpdateStudent(req *domain.Student) (*domain.Student, error)
	DeleteStudent(id int) (int, error)
}

type StudentHandler struct {
	svc StudentService
}

func NewStudentHandler(svc StudentService) *StudentHandler {
	return &StudentHandler{svc: svc}
}

func (h *StudentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var s domain.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if s.FirstName == "" || s.LastName == "" || s.Email == "" {
		writeJSONError(w, "missing required fields", http.StatusBadRequest)
		return
	}

	createdStudent, err := h.svc.CreateStudent(&s)
	if err != nil {
		writeJSONError(w, "failed to create student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(createdStudent); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *StudentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid student ID", http.StatusBadRequest)
		return
	}

	s, err := h.svc.GetStudentByID(id)
	if err != nil {
		writeJSONError(w, "student not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *StudentHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	students, err := h.svc.GetAllStudents()
	if err != nil {
		writeJSONError(w, "failed to fetch students", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(students); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *StudentHandler) Update(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid student ID", http.StatusBadRequest)
		return
	}

	var s domain.Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}
	s.ID = id

	updatedStudent, err := h.svc.UpdateStudent(&s)
	if err != nil {
		writeJSONError(w, "failed to update student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updatedStudent); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}

func (h *StudentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, "invalid student ID", http.StatusBadRequest)
		return
	}

	deletedID, err := h.svc.DeleteStudent(id)
	if err != nil {
		writeJSONError(w, "failed to delete student", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]int{"deleted_id": deletedID}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		slog.Error("error encoding response", "error", err)
	}
}
