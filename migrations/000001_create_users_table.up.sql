CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    fullname VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    gender VARCHAR(10) NOT NULL,
    birthdate DATE NOT NULL,
    image VARCHAR(255),
    interests VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    is_first_login BOOLEAN DEFAULT FALSE,
    password VARCHAR(255) NOT NULL,
    token VARCHAR(255),
    user_type_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)