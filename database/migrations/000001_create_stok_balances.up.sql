CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE "stok_balances" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    barang_id UUID NOT NULL,
    lokasi_id UUID NOT NULL,
    qty_balance integer DEFAULT 0,
    isActive boolean DEFAULT true,
    created_by UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
)

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_stok_balances_updated_at 
BEFORE UPDATE ON stok_balances
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
