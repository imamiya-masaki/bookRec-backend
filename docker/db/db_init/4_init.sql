/* */
INSERT INTO book_db.users (username, twitter_token) VALUES 
("A木　太郎", "dwwadamlkdfmwaklrakmralrma"), ("B木　太郎", "twitterid2"), ("C木　太郎", "twitterid3"), ("D木　太郎", "twitterid4"), ("E木　太郎", "twitterid5");
INSERT INTO book_db.recommends (sender_id, receiver_id, book_id, reaction_content_id) VALUES
(1,2, 1, 0),(1,3,1,0),(2,3,2,0),(3,4,2,0),(4,5,3,0),(5,1,3,0);
INSERT INTO book_db.messages (reccomend_id, sendedtime, content) VALUES
(1,NULL,"dada"),(2,NULL,"dada"),(3,NULL,"dada");
INSERT INTO book_db.book_group (bookname) VALUES
("きめつ"),("なると");
INSERT INTO book_db.books (book_group_id, title, author, price, uri) VALUES
(1,"きめつ(1)","吾峠呼世晴",1000,"https://res.cloudinary.com/teamb/image/upload/v1600313784/bj01_pages-to-jpg-0002_vd5k8r.jpg"),
(1,"きめつ(2)","吾峠呼世晴",1000, "https://res.cloudinary.com/teamb/image/upload/v1600390594/%E3%83%95%E3%82%99%E3%83%A9%E3%83%83%E3%82%AF%E3%82%B7%E3%82%99%E3%82%A7%E3%83%83%E3%82%AF_ehjhvm.jpg"),
(2,"なると(10)","岸本 斉史",3000, "https://res.cloudinary.com/teamb/image/upload/v1600390603/%E3%83%95%E3%82%99%E3%83%A9%E3%83%83%E3%82%AF%E3%82%B7%E3%82%99%E3%83%A3%E3%83%83%E3%82%AF%E3%81%AB%E3%82%88%E3%82%8D%E3%81%97%E3%81%8F3_mtepjc.jpg");
INSERT INTO book_db.my_books (user_id, book_id) VALUES
(1,1),(1,2),(1,3),(2,1),(2,2),(3,3);
INSERT INTO book_db.booklist (user_id, book_id) VALUES
(1,1),(1,2),(1,3);
INSERT INTO book_db.reaction_contents (reaction_name, uri) VALUES
("自己啓発", "https://res.cloudinary.com/teamb/image/upload/v1600394098/%E8%87%AA%E5%B7%B1%E5%95%93%E7%99%BA_yb7yeb.png");
INSERT INTO book_db.coupons (user_id, percent) VALUES (1, 10);
