DROP TABLE roles;
CREATE TABLE roles
(
    id         uuid      NOT NULL PRIMARY KEY UNIQUE DEFAULT uuid_generate_v4(),
    user_id    uuid      NOT NULL UNIQUE REFERENCES users ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL                    DEFAULT current_timestamp,
    admin      BOOLEAN   NOT NULL                    default false,
    member     BOOLEAN   NOT NULL                    default false
);
