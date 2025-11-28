CREATE TABLE IF NOT EXISTS donation_receipts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    muzakki_id UUID NOT NULL REFERENCES muzakki(id) ON DELETE RESTRICT,
    receipt_number VARCHAR(50) NOT NULL UNIQUE,
    receipt_date DATE NOT NULL,
    payment_method VARCHAR(50) NOT NULL,
    total_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    notes TEXT,
    created_by_user_id UUID NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_donation_receipts_receipt_number ON donation_receipts(receipt_number);
CREATE INDEX IF NOT EXISTS idx_donation_receipts_receipt_date ON donation_receipts(receipt_date);
CREATE INDEX IF NOT EXISTS idx_donation_receipts_muzakki_id ON donation_receipts(muzakki_id);
CREATE INDEX IF NOT EXISTS idx_donation_receipts_created_by_user_id ON donation_receipts(created_by_user_id);
