CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id              uuid not null unique default uuid_generate_v4(),
    username        varchar unique,
    first_name      varchar,
    last_name       varchar,
    email           varchar unique,
    phone           varchar,
    hashed_password varchar,
    password_salt   varchar,
    created_at      timestamp not null default current_timestamp,
    last_updated    timestamp not null default current_timestamp
);