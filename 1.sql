CREATE DATABASE miniproject charset=UTF8;

USE miniproject;

CREATE TABLE users(
    id VARCHAR(10) PRIMARY KEY,
    nickname VARCHAR(20),
    avatar TEXT,
    buygoods TEXT
);

CREATE TABLE goods(
    goods_id INT PRIMARY KEY AUTO_INCREMENT,
    id VARCHAR(10),
    title TEXT,
    price INT ,
    goodszone TEXT,
    summary TEXT,
    scores FLOAT,
    goodsin VARCHAR(5) DEFAULT "yes",
    way TEXT,
    avatar TEXT,
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
    comment TEXT,
    givetime TEXT,
    CONSTRAINT third
    FOREIGN KEY(id) REFERENCES users(id),
    CONSTRAINT forth
    FOREIGN KEY(goods_id) REFERENCES goods(goods_id)
);

CREATE TABLE messages(
    id INT PRIMARY KEY AUTO_INCREMENT,
    buyer VARCHAR(10),
    my VARCHAR(10),
    msg TEXT
)