CREATE DATABASE miniproject charset=UTF8;

USE miniproject;

CREATE TABLE users(
    id VARCHAR(10) PRIMARY KEY,
    nickname VARCHAR(20),
    avatar TEXT
);

CREATE TABLE goods(
    goodsid INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    scores INT,
    title TEXT,
    price INT NOT NULL,
    goodszone TEXT,
    summary TEXT,
    way TEXT,
    avatar TEXT,
    buyer TEXT DEFAULT "",
    feedback INT DEFAULT 0,
    CONSTRAINT first
    FOREIGN KEY(id) REFERENCES user(id)
);

CREATE TABLE carts(
    id VARCHAR(10) PRIMARY KEY,
    goodsid DEFAULT "",
    CONSTRAINT second
    FOREIGN KEY(id) REFERENCES user(id)
);

CREATE TABLE comments(
    commentid INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    score INT,
    goodsid INT,
    comment TEXT,
    CONSTRAINT third
    FOREIGN KEY(id) REFERENCES user(id),
    CONSTRAINT forth
    FOREIGN KEY(goodsid) REFERENCES good(goodsid)
);