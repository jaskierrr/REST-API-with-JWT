BEGIN;

ALTER TABLE users DROP COLUMN name;
ALTER TABLE users ADD COLUMN email VARCHAR(320);
ALTER TABLE users ADD COLUMN password_hash TEXT;

UPDATE users
SET email = 'user' || id || '@example.com',
    password_hash = 'placeholder_hash_' || id;

ALTER TABLE users ALTER COLUMN email SET NOT NULL;
ALTER TABLE users ALTER COLUMN password_hash SET NOT NULL;
ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE (email);
ALTER TABLE users ADD CONSTRAINT unique_password_hash UNIQUE (password_hash);

COMMIT;
