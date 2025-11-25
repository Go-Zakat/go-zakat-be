CREATE TABLE IF NOT EXISTS muzakki (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    phoneNumber VARCHAR(255) UNIQUE NOT NULL,
    address VARCHAR(255),
    notes VARCHAR(255),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_muzakki_name ON muzakki(name);