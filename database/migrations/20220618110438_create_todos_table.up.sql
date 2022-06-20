CREATE TABLE IF NOT EXISTS todos(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    title VARCHAR NOT NULL,
    description TEXT,
    created_at TIME DEFAULT now()
);