ALTER TABLE roles ADD COLUMN admin BOOLEAN;
ALTER TABLE roles ADD COLUMN member BOOLEAN;
ALTER TABLE roles DROP COLUMN type;