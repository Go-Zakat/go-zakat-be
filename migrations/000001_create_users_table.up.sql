CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(), -- id unik
    email VARCHAR(255) UNIQUE NOT NULL,            -- email, harus unik
    password VARCHAR(255),                         -- password hash (boleh NULL untuk user Google only)
    google_id VARCHAR(255) UNIQUE,                 -- untuk login via Google
    name VARCHAR(255) NOT NULL,                    -- nama user
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- kapan dibuat
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()  -- kapan diupdate
);

CREATE INDEX idx_users_email ON users(email);
