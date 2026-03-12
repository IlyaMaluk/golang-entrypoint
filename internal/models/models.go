package models

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Teacher struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Department string `json:"department"`
}

type Course struct {
	ID          int
	Title       string
	Description string
	TeacherID   int
}

type Enrollment struct {
	StudentID int `json:"student_id"`
	CourseID  int `json:"course_id"`
}
