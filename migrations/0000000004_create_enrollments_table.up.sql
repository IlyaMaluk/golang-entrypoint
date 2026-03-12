CREATE TABLE enrollments
(
    student_id INT NOT NULL REFERENCES students (id) ON DELETE CASCADE,
    course_id  INT NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    PRIMARY KEY (student_id, course_id)
);