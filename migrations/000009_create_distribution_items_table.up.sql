CREATE TABLE IF NOT EXISTS distribution_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    distribution_id UUID NOT NULL REFERENCES distributions(id) ON DELETE CASCADE,
    mustahiq_id UUID NOT NULL REFERENCES mustahiq(id) ON DELETE RESTRICT,
    amount DECIMAL(15, 2) NOT NULL,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_distribution_items_distribution_id ON distribution_items(distribution_id);
CREATE INDEX IF NOT EXISTS idx_distribution_items_mustahiq_id ON distribution_items(mustahiq_id);
