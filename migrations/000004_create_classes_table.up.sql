CREATE TABLE classes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255),
    thumbnail VARCHAR(255),
    description VARCHAR(255),
    meta_description VARCHAR(255),
    level VARCHAR(10) NOT NULL,
    class_category_id INT NOT NULL,
    tags VARCHAR(255),
    slug VARCHAR(255) NOT NULL,
    method VARCHAR(255),
    media VARCHAR(255),
    prefix_code VARCHAR(255),
    materials VARCHAR(255),
    collaboration_feed VARCHAR(255),
    instructor_id INT,
    learning_link VARCHAR(255),
    consultancy_link VARCHAR(255),
    consultancy_schedule VARCHAR(255),
    group_chat_link VARCHAR(255),
    registration_close_date DATE,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
)