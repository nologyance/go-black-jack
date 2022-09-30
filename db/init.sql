CREATE USER jack;

CREATE DATABASE jack;

GRANT ALL PRIVILEGES ON DATABASE jack to jack;

\c jack

CREATE TABLE player (
    id integer,
    name varchar(30)
);

INSERT INTO player VALUES (1, 'namiki');