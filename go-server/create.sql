CREATE TABLE Users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL UNIQUE,
    passwrd VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    gameid INT
);

SELECT * FROM Users;

CREATE TABLE Tickets (
    id INT PRIMARY KEY AUTO_INCREMENT,
    usersid INT NOT NULL,
    placeid INT NOT NULL
);

CREATE TABLE Places (
    id INT PRIMARY KEY AUTO_INCREMENT,
    place VARCHAR(255) NOT NULL,
    day DATETIME NOT NULL UNIQUE,
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