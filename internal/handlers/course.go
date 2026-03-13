package handlers

import (
	"context"
	"encoding/json"
	"golang-entrypoint/internal/service"
	"log/slog"
	"net/http"
	"strconv"

	"golang-entrypoint/internal/domain"
)

type CourseService interface {
	CreateCourse(ctx context.Context, req *domain.Course) (*domain.Course, error)
	GetCourseByID(ctx context.Context, id int) (*domain.Course, error)
	GetAllCourses(ctx context.Context) ([]domain.Course, error)
	UpdateCourse(ctx context.Context, req *domain.Course) (*domain.Course, error)
	DeleteCourse(ctx context.Context, id int) (int, error)
}

type CourseHandler struct {
	svc              CourseService
	validatorService service.ValidatorService
}

func NewCourseHandler(
	svc CourseService,
	validatorService service.ValidatorService,
) *CourseHandler {
	return &CourseHandler{
		svc:              svc,
		validatorService: validatorService,
	}
}

func (h *CourseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c domain.Course
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		writeJSONError(w, "invalid request body", http.StatusBadRequest)
		return
	}

	validationErrors, err := h.validatorService.Validate(&c)
	if err != nil {
		writeJSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(validationErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validationErrors)
		return
	}

	ctx := r.Context()
	createdCourse, err := h.svc.CreateCourse(ctx, &c)
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

	ctx := r.Context()
	c, err := h.svc.GetCourseByID(ctx, id)
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

func (h *CourseHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	courses, err := h.svc.GetAllCourses(ctx)
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

	ctx := r.Context()
	updatedCourse, err := h.svc.UpdateCourse(ctx, &c)
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

	ctx := r.Context()
	deletedID, err := h.svc.DeleteCourse(ctx, id)
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
