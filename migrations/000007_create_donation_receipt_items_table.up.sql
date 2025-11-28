CREATE TABLE IF NOT EXISTS donation_receipt_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    receipt_id UUID NOT NULL REFERENCES donation_receipts(id) ON DELETE CASCADE,
    fund_type VARCHAR(20) NOT NULL CHECK (fund_type IN ('zakat', 'infaq', 'sadaqah')),
    zakat_type VARCHAR(20) CHECK (zakat_type IN ('fitrah', 'maal')),
    person_count INT,
    amount DECIMAL(15, 2) NOT NULL,
    rice_kg DECIMAL(10, 2),
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_donation_receipt_items_receipt_id ON donation_receipt_items(receipt_id);
CREATE INDEX IF NOT EXISTS idx_donation_receipt_items_fund_type ON donation_receipt_items(fund_type);
