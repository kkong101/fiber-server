create table TEST_USER
(
    idx      int auto_increment
        primary key,
    name     varchar(100) not null,
    phone    varchar(20)  not null,
    birthday datetime     not null,
    password varchar(200) null
);

