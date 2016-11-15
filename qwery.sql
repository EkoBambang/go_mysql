MariaDB [(none)]> create database praweb;
Query OK, 1 row affected (0.08 sec)

MariaDB [(none)]> use praweb;
Database changed
MariaDB [praweb]> create table mhs(
id_mhs char(8),
nama_mhs varchar(25));

insert into mhs values
(101,'eko'),
(102,'bambang'),
(103,'sinu'),
(104,'darsono'),
(105,'ahmad'),
(106,'deniza'),
(107,'tirtana'),
(108,'moch'),
(109,'soni'),
(110,'casini');

