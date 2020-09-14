CREATE TABLE users (
    id INT auto_increment not null primary key,
    username VARCHAR(32),
    twitterToken VARCHAR(256)
);

CREATE TABLE recommends (
    id INT auto_increment not null primary key,
    senderId INT,
    receiverId INT,
    bookId INT,
    reactionContentId INT
);

CREATE TABLE messages (
    id INT auto_increment not null primary key,
    reccomendId INT,
    sendedtime TIMESTAMP,
    content VARCHAR(256)
);

CREATE TABLE booklist (
    id INT auto_increment not null primary key,
    userId INT,
    bookId INT
);

CREATE TABLE myBooks (
    id INT auto_increment not null primary key,
    userId INT,
    bookId INT
);

CREATE TABLE bookGroup (
    id INT auto_increment not null primary key,
    bookname VARCHAR(64)
);

CREATE TABLE books (
    id INT auto_increment not null primary key,
    bookGroupId INT,
    title VARCHAR(64),
    author VARCHAR(64),
    price INT,
    releaseDate VARCHAR(128)
);

CREATE TABLE bookContents (
    id INT auto_increment not null primary key,
    bookId INT,
    bookPage INT,
    URI VARCHAR(256)
);
CREATE TABLE reactionContent (
    id INT auto_increment not null primary key,
    reactionName VARCHAR(64),
    URI VARCHAR(256)
);