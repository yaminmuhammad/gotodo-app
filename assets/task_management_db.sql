CREATE DATABASE task_management_db;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE authors (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50),
    email VARCHAR(100) unique,
    password VARCHAR(100),
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp
);

CREATE TABLE tasks (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    title VARCHAR(100),
    content TEXT,
    author_id uuid,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp,
    FOREIGN KEY(author_id) REFERENCES authors(id)
);

INSERT INTO authors (name, email, password) VALUES
('John', 'john@gmail.com', 'password'),
('Tailor', 'tailor@gmail.com', 'password');

INSERT INTO tasks (title, content, author_id) VALUES
('Golang', 'Belajar Golang with GIN', '7acbafb2-56c5-4a1f-bd9d-05302e4db5b9'),
('Golang', 'Belajar Golang with GIN', '4a223243-72e4-4bc1-af27-ee2b94d84142');