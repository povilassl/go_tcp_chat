create table users (
    id varchar(36) not null primary key,
    username varchar(32) not null unique,
    password_hash varchar(255) not null,
    created_at timestamp not null default current_timestamp
)