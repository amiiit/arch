ALTER TABLE roles
    DROP COLUMN admin;
ALTER TABLE roles
    DROP COLUMN member;
ALTER TABLE roles
    ADD COLUMN type varchar;