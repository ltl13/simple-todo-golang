create database todo_manager;
use todo_manager;
create table TodoItem (
    id int auto_increment,
    title varchar(100),
    detail varchar(200),
    isdone tinyint,
    primary key(id)
);