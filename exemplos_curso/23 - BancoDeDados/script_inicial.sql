CREATE DATABASE devbook;

USE devbook;

create table usuarios (
    id int AUTO_INCREMENT not null PRIMARY key,
    nome varchar(50) not null,
    email varchar(50) not null
) ENGINE=INNODB;

CREATE USER 'golang'@'localhost' IDENTIFIED BY 'goland'
GRANT ALL privileges ON devbook.* TO 'golang'@'localhost'