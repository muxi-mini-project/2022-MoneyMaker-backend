CREATE DATABASE miniproject charset=UTF8;

USE miniproject;

CREATE TABLE user(
    id VARCHAR(10) PRIMARY KEY,
    nickname VARCHAR(20),
    password VARCHAR(15) NOT NULL,
    avatar TEXT
);

CREATE TABLE good(
    goodsid INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    scores INT,
    price INT NOT NULL,
    zone TEXT,
    summary TEXT,
    way TEXT,
    buyer TEXT,
    feedback INT,
    in VARCHAR(3),
    CONSTRAINT first
    FOREIGN KEY(id) REFERENCES user(id)
);

CREATE TABLE cart(
    id VARCHAR(10),
    owner TEXT,
    CONSTRAINT second
    FOREIGN KEY(id) REFERENCES user(id)
);

CREATE TABLE comment(
    id VARCHAR(10),
    score INT,
    goodsid INT,
    comment TEXT,
    CONSTRAINT third
    FOREIGN KEY(id) REFERENCES user(id),
    CONSTRAINT forth
    FOREIGN KEY(goodsid) REFERENCES good(goodsid)
);