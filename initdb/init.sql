CREATE DATABASE IF NOT EXISTS shortenurl;
CREATE USER 'trangx'@'localhost' IDENTIFIED BY '';
GRANT ALL PRIVILEGES ON shortenurl.* TO 'trangx'@'localhost';

use shortenurl;
create table IF NOT EXISTS test (
    full varchar(50),
    short varchar(50)
);