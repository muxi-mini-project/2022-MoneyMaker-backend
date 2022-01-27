DROP DATABASE if exists miniproject;

CREATE DATABASE miniproject charset=UTF8;

USE miniproject;

CREATE TABLE users(
    id VARCHAR(10) PRIMARY KEY,
    nickname VARCHAR(20),
    avatar VARCHAR(100),
    buygoods TEXT
);

CREATE TABLE goods(
    goods_id INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    title VARCHAR(50),
    price INT ,
    goodszone VARCHAR(50),
    summary VARCHAR(100),
    scores FLOAT,
    goodsin VARCHAR(5) DEFAULT "yes",
    way VARCHAR(100),
    avatar VARCHAR(100),
    buyer TEXT ,
    feed_back INT DEFAULT 0,
    CONSTRAINT first
    FOREIGN KEY(id) REFERENCES users(id)
);

CREATE TABLE carts(
    id VARCHAR(10) PRIMARY KEY,
    goodsid TEXT,
    CONSTRAINT second
    FOREIGN KEY(id) REFERENCES users(id)
);

CREATE TABLE comments(
    comment_id INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    score INT,
    goods_id INT,
    comment VARCHAR(100),
    givetime VARCHAR(20),
    CONSTRAINT third
    FOREIGN KEY(id) REFERENCES users(id),
    CONSTRAINT forth
    FOREIGN KEY(goods_id) REFERENCES goods(goods_id)
);

CREATE TABLE messages(
    id INT PRIMARY KEY AUTO_INCREMENT,
    buyer VARCHAR(10),
    my VARCHAR(10),
    msg VARCHAR(30)
)