CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users
(
    id              uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    username        VARCHAR UNIQUE,
    first_name      VARCHAR,
    last_name       VARCHAR,
    email           VARCHAR UNIQUE,
    phone           VARCHAR,
    region          VARCHAR,
    hashed_password VARCHAR,
    created_at      TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    last_updated    TIMESTAMP NOT NULL                    DEFAULT current_timestamp
);


CREATE TABLE sessions
(
    id         uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    user_id    uuid      NOT NULL REFERENCES users ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    token      varchar   NOT NULL UNIQUE,
    is_valid   boolean   NOT NULL                    DEFAULT true
);
CREATE UNIQUE INDEX IF NOT EXISTS sessions_token ON sessions (token);

CREATE TABLE roles
(
    id         uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    user_id    uuid      NOT NULL REFERENCES users ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    type       varchar   NOT NULL,
    UNIQUE (user_id, type)
)
