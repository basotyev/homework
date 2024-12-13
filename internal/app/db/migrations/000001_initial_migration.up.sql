CREATE TABLE users(
    id serial primary key,
    name varchar(32),
    email varchar(64),
    age integer,
    created_at timestamp default now(),
    updated_at timestamp default now()
);
