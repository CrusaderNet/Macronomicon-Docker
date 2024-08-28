
-- Create a new database
-- CREATE DATABASE macronomicon;

--mysql macronomicon --password


-- Create User Table
CREATE TABLE users(
    id INT,
    email VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

-- Create Macro Entries Table
CREATE TABLE macro_entries(
    id INT,
    user_id INT,
    submit_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP(),
    proteins FLOAT NOT NULL,
    carbs FLOAT NOT NULL,
    fats FLOAT NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

--Fake Data test
INSERT INTO users (id, email) VALUES (1, 'stourish@comcast.net');
INSERT INTO users (id, email) VALUES (2, 'Nyko@gmail.com');

INSERT INTO macro_entries (id, user_id, submit_time, proteins, carbs, fats) 
VALUES (1, 1, CURRENT_TIMESTAMP(), 10.6, 3.5, 7.8);

INSERT INTO macro_entries (id, user_id, submit_time, proteins, carbs, fats) 
VALUES (2, 2, CURRENT_TIMESTAMP(), 25.6, 7.5, 6.8);

INSERT INTO macro_entries (id, user_id, submit_time, proteins, carbs, fats) 
VALUES (3, 2, CURRENT_TIMESTAMP(), 35.6, 4.5, 9.8);