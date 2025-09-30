CREATE TABLE users (
    id CHAR(36) PRIMARY KEY, -- UUID armazenado como string (36 caracteres)
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL, -- guarda o hash da senha (bcrypt gera até 60 chars)
    creation_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_date DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (id, name, email, password, creation_date, update_date)
VALUES (
    UUID(),
    'Douglas Junior',
    'douglasJunior@example.com',
    '$2a$10$twK2.FqHkYFkF0mEGOZWCOdJ3tr.T2C8McMlxsZHhJF1ul2LCduri', --123456
    NOW(),
    NOW()
);