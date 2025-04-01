CREATE TABLE IF NOT EXISTS people (
                                      id SERIAL PRIMARY KEY,
                                      name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255),
    age INTEGER NOT NULL,
    gender VARCHAR(50) NOT NULL,
    nationality VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );