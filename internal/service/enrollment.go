package service

type EnrollmentRepository interface {
	EnrollStudent(studentID, courseID int) error
	UnenrollStudent(studentID, courseID int) error
}

type EnrollmentServiceImpl struct {
	repo EnrollmentRepository
}

func NewEnrollmentService(repo EnrollmentRepository) *EnrollmentServiceImpl {
	return &EnrollmentServiceImpl{repo: repo}
}

func (s *EnrollmentServiceImpl) Enroll(studentID, courseID int) error {
	return s.repo.EnrollStudent(studentID, courseID)
}

func (s *EnrollmentServiceImpl) Unenroll(studentID, courseID int) error {
	return s.repo.UnenrollStudent(studentID, courseID)
}
