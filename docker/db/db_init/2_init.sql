CREATE TABLE posts (
    id int auto_increment not null primary key, 
    title varchar(32), 
    content varchar(256)
);

insert into book_db.posts (title, content) values ('hoge', 'hogehoge');
