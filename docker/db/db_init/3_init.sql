CREATE TABLE users (
    id INT auto_increment not null primary key,
    username VARCHAR(32),
    twitter_token VARCHAR(256)
);

CREATE TABLE recommends (
    id INT auto_increment not null primary key,
    sender_id INT,
    receiver_id INT,
    book_id INT,
    reaction_content_id INT
);

CREATE TABLE messages (
    id INT auto_increment not null primary key,
    reccomend_id INT,
    sendedtime TIMESTAMP,
    content VARCHAR(256)
);

CREATE TABLE booklist (
    id INT auto_increment not null primary key,
    user_id INT,
    book_id INT
);

CREATE TABLE myBooks (
    id INT auto_increment not null primary key,
    user_id INT,
    book_id INT
);

CREATE TABLE bookGroup (
    id INT auto_increment not null primary key,
    bookname VARCHAR(64)
);

CREATE TABLE books (
    id INT auto_increment not null primary key,
    book_group_id INT,
    title VARCHAR(64),
    author VARCHAR(64),
    price INT,
    release_date VARCHAR(128)
);

CREATE TABLE bookContents (
    id INT auto_increment not null primary key,
    book_id INT,
    book_page INT,
    uri VARCHAR(256)
);
CREATE TABLE reactionContent (
    id INT auto_increment not null primary key,
    reaction_name VARCHAR(64),
    uri VARCHAR(256)
);