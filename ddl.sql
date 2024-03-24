CREATE DATABASE employees_data; -- for local test

CREATE TABLE employees (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  phone VARCHAR(20) NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);

-- Data Dummy
INSERT INTO employees (name, email, phone)
VALUES ('Dani Santoso', 'budi@mail.com', '08123456789'),
       ('Ani Puspita', 'ani.puspita@mail.com', '08987654321'),
       ('Awen Setiawan', 'awen.setiawan@mail.com', '08574123456'),
       ('Fajar Purnama', 'fajar.purnama@mail.com', '0875551234'),
       ('Safar Wijaya', 's.wijaya@mail.com', '08219876543');


