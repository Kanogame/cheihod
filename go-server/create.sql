CREATE TABLE Users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    passwrd VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    gameid INT
);

DROP TABLE Users;

SELECT * FROM Users;

CREATE TABLE Tickets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    usersid INT NOT NULL,
    placeid INT NOT NULL
);

CREATE TABLE Places (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    place VARCHAR(255) NOT NULL,
    time DATETIME NOT NULL UNIQUE,
    capacity INT NOT NULL,
    deprecated BOOLEAN NOT NULL DEFAULT(false)
);

DROP TABLE Places;

CREATE TABLE Games (
    id INT PRIMARY KEY AUTO_INCREMENT,
    gamename VARCHAR(255) NOT NULL,
    gameToken VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE Admins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    passwrd VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE
);

CREATE TABLE Token (
    id INT PRIMARY KEY AUTO_INCREMENT,
    token VARCHAR(30) NOT NULL UNIQUE,
    userid INT NOT NULL
);

SELECT * FROM Tickets;

SELECT name, place, time, capacity FROM Places WHERE time BETWEEN date_sub(now(), interval 1 MONTH) AND date_add(now(), interval 1 MONTH);

SELECT name, place, time FROM Places WHERE id IN (SELECT placeid FROM Tickets WHERE usersid = 1)
 AND time BETWEEN current_timestamp() AND date_add(now(), interval 1 MONTH);