CREATE TABLE IF NOT EXISTS programs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    type VARCHAR(100) NOT NULL,
    description TEXT,
    active BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_programs_name ON programs(name);
CREATE INDEX IF NOT EXISTS idx_programs_type ON programs(type);
CREATE INDEX IF NOT EXISTS idx_programs_active ON programs(active);
