package service

import "context"

type EnrollmentRepository interface {
	EnrollStudent(ctx context.Context, studentID, courseID int) error
	UnenrollStudent(ctx context.Context, studentID, courseID int) error
}

type EnrollmentServiceImpl struct {
	repo EnrollmentRepository
}

func NewEnrollmentService(repo EnrollmentRepository) *EnrollmentServiceImpl {
	return &EnrollmentServiceImpl{repo: repo}
}

func (s *EnrollmentServiceImpl) Enroll(ctx context.Context, studentID, courseID int) error {
	return s.repo.EnrollStudent(ctx, studentID, courseID)
}

func (s *EnrollmentServiceImpl) Unenroll(ctx context.Context, studentID, courseID int) error {
	return s.repo.UnenrollStudent(ctx, studentID, courseID)
}
