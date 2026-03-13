package domain

type Course struct {
	ID          int    `json:"id"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	TeacherID   int    `json:"teacher_id" validate:"required"`
}
