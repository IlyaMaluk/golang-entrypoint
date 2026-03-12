CREATE TABLE teachers
(
    id         SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name  VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL
);