package main

import (
	"context"
	"errors"
	"golang-entrypoint/internal/service"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang-entrypoint/internal/database"
	"golang-entrypoint/internal/handlers"
	"golang-entrypoint/internal/repository"
	_ "golang-entrypoint/internal/service"
	"golang-entrypoint/pkg/middleware"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "postgres://user:password@localhost:5432/school_db?sslmode=disable"
	}

	db, err := database.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	slog.Info("Migrations applied")

	studentRepo := repository.NewStudentRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	courseRepo := repository.NewCourseRepository(db)
	enrollmentRepo := repository.NewEnrollmentRepository(db)

	courseService := service.NewCourseService(courseRepo)

	studentHandler := handlers.NewStudentHandler(studentRepo)
	teacherHandler := handlers.NewTeacherHandler(teacherRepo)
	courseHandler := handlers.NewCourseHandler(courseService)
	enrollmentHandler := handlers.NewEnrollmentHandler(enrollmentRepo)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", handlers.Login)

	mux.HandleFunc("GET /students", studentHandler.GetAll)
	mux.HandleFunc("GET /students/{id}", studentHandler.GetByID)
	mux.HandleFunc("POST /students", middleware.JWTAuth(studentHandler.Create))
	mux.HandleFunc("PUT /students/{id}", middleware.JWTAuth(studentHandler.Update))
	mux.HandleFunc("DELETE /students/{id}", middleware.JWTAuth(studentHandler.Delete))

	mux.HandleFunc("GET /teachers", teacherHandler.GetAll)
	mux.HandleFunc("GET /teachers/{id}", teacherHandler.GetByID)
	mux.HandleFunc("POST /teachers", middleware.JWTAuth(teacherHandler.Create))
	mux.HandleFunc("PUT /teachers/{id}", middleware.JWTAuth(teacherHandler.Update))
	mux.HandleFunc("DELETE /teachers/{id}", middleware.JWTAuth(teacherHandler.Delete))

	mux.HandleFunc("GET /courses", courseHandler.GetAll)
	mux.HandleFunc("GET /courses/{id}", courseHandler.GetByID)
	mux.HandleFunc("POST /courses", middleware.JWTAuth(courseHandler.Create))
	mux.HandleFunc("PUT /courses/{id}", middleware.JWTAuth(courseHandler.Update))
	mux.HandleFunc("DELETE /courses/{id}", middleware.JWTAuth(courseHandler.Delete))

	mux.HandleFunc("POST /students/{id}/courses/{course_id}", middleware.JWTAuth(enrollmentHandler.Enroll))
	mux.HandleFunc("DELETE /students/{id}/courses/{course_id}", middleware.JWTAuth(enrollmentHandler.Unenroll))

	loggedMux := middleware.Logger(mux)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggedMux,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	slog.Info("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown Failed: %+v", err)
	}

	slog.Info("Server exited properly")
}
