USE students_database;

CREATE TABLE students
(
    id      INT AUTO_INCREMENT PRIMARY KEY,
    name    VARCHAR(100),
    age     INT,
    gender  VARCHAR(10),
    grade   VARCHAR(10),
    address VARCHAR(100)
);

INSERT INTO students (name, age, gender, grade, address)
VALUES ('John', 18, 'Male', 'A', '123 Main St'),
       ('Jane', 19, 'Female', 'B', '456 Elm St'),
       ('Mike', 20, 'Male', 'B', '789 Oak St');