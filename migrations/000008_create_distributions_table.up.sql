CREATE TABLE IF NOT EXISTS distributions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    distribution_date DATE NOT NULL,
    program_id UUID REFERENCES programs(id) ON DELETE RESTRICT,
    source_fund_type VARCHAR(20) NOT NULL CHECK (source_fund_type IN ('zakat_fitrah', 'zakat_maal', 'infaq', 'sadaqah')),
    total_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    notes TEXT,
    created_by_user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_distributions_distribution_date ON distributions(distribution_date);
CREATE INDEX IF NOT EXISTS idx_distributions_program_id ON distributions(program_id);
CREATE INDEX IF NOT EXISTS idx_distributions_source_fund_type ON distributions(source_fund_type);
CREATE INDEX IF NOT EXISTS idx_distributions_created_by_user_id ON distributions(created_by_user_id);
