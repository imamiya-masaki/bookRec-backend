/* */
INSERT INTO book_db.users (username, twitter_token) VALUES 
("A木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("B木　太郎", "twitterid2"), ("C木　太郎", "twitterid3"), ("D木　太郎", "twitterid4"), ("E木　太郎", "twitterid5");
INSERT INTO book_db.recommends (sender_id, receiver_id, book_id, reaction_content_id) VALUES
(1,2, 1, 0),(1,3,1,0),(2,3,2,0),(3,4,2,0),(4,5,3,0),(5,1,3,0);
INSERT INTO book_db.messages (reccomend_id, sendedtime, content) VALUES
(1,NULL,"dada"),(2,NULL,"dada"),(3,NULL,"dada");
INSERT INTO book_db.book_group (bookname) VALUES
("きめつ"),("なると");
INSERT INTO book_db.books (book_group_id, title, author, price) VALUES
(1,"きめつ(1)","吾峠呼世晴",1000),(1,"きめつ(2)","吾峠呼世晴",1000),(2,"なると(10)","岸本 斉史",3000);
INSERT INTO book_db.my_books (user_id, book_id) VALUES
(1,1),(1,2),(1,3),(2,1),(2,2),(3,3);
INSERT INTO book_db.booklist (user_id, book_id) VALUES
(1,1),(1,2),(1,3);
INSERT INTO book_db.reaction_contents (reaction_name, uri) VALUES
(NULL,"hogehoge");

INSERT INTO book_db.coupons (user_id, percent) VALUES (1, 10);

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

