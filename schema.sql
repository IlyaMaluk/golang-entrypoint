CREATE TABLE teachers
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL
);

CREATE TABLE students
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100)        NOT NULL,
    last_name  VARCHAR(100)        NOT NULL,
    email      VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE courses
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    teacher_id  INT          NOT NULL REFERENCES teachers (id) ON DELETE CASCADE
);

CREATE TABLE enrollments
(
    student_id INT NOT NULL REFERENCES students (id) ON DELETE CASCADE,
    course_id  INT NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    PRIMARY KEY (student_id, course_id)
);