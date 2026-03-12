CREATE TABLE courses
(
    id          SERIAL PRIMARY KEY,
    title       VARCHAR(255) NOT NULL,
    description TEXT,
    teacher_id  INT          NOT NULL REFERENCES teachers (id) ON DELETE CASCADE
);
