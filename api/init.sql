--mysql macronomicon --password
CREATE DATABASE macronomicon;

USE macronomicon;

CREATE USER 'user'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON *.* TO 'user'@'%';
FLUSH PRIVILEGES;

-- Create User Table
CREATE TABLE users(
    id INT,
    email VARCHAR(256) NOT NULL,
    PRIMARY KEY (id)
);

-- Create Macro Entries Table
CREATE TABLE macro_entries(
    id INT NOT NULL AUTO_INCREMENT,
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

INSERT INTO macro_entries (user_id, submit_time, proteins, carbs, fats) 
VALUES (1, CURRENT_TIMESTAMP(), 10.6, 3.5, 7.8);

INSERT INTO macro_entries (user_id, submit_time, proteins, carbs, fats) 
VALUES (2, CURRENT_TIMESTAMP(), 25.6, 7.5, 6.8);

INSERT INTO macro_entries (user_id, submit_time, proteins, carbs, fats) 
VALUES (2, CURRENT_TIMESTAMP(), 35.6, 4.5, 9.8);