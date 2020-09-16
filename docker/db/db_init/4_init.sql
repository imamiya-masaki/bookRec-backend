/* */
INSERT INTO book_db.users (username, twitter_token) VALUES 
("A木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("B木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("C木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("D木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("E木　太郎", "dwwadamlkdfmwaklrakmralrma");
INSERT INTO book_db.recommends (sender_id, receiver_id, book_id, reaction_content_id) VALUES
(1,2, 1, 1),(1,3,1,1),(2,3,2,1);
INSERT INTO book_db.messages (reccomend_id, sendedtime, content) VALUES
(1,NULL,"dada"),(2,NULL,"dada"),(3,NULL,"dada");
INSERT INTO book_db.book_group (bookname) VALUES
("毀滅のヤイバ!!"),("なると");
INSERT INTO book_db.books (book_group_id, title, author, price, release_date) VALUES
(1,"毀滅のヤイバ！！(1)","吾峠呼世晴",3000,NULL),(1,"毀滅のヤイバ！！(2)","吾峠呼世晴",3000,NULL),(2,"なると(10)","岸本 斉史",3000,NULL);
INSERT INTO book_db.my_books (user_id, book_id) VALUES
(1,1),(2,1),(1,2),(1,3);
INSERT INTO book_db.booklist (user_id, book_id) VALUES
(1,1),(1,2),(1,3);
INSERT INTO book_db.book_contents (book_id, book_page, uri) VALUES
(1,1,NULL),(1,2,NULL),(1,3,NULL),(1,4,NULL),(1,5,NULL),(1,6,NULL);
INSERT INTO book_db.reactionContent (reaction_name, uri) VALUES
(NULL,NULL);
/*
CREATE TABLE booklist (
    id INT auto_increment not null primary key,
    user_id INT,
    book_id INT
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
*/

