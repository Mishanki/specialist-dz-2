CREATE TABLE users (
    id serial PRIMARY KEY,
    username varchar not null unique,
    password varchar not null
);