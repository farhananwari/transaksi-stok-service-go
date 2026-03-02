
CREATE TABLE "transaksi_stoks" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    no_transaksi VARCHAR(100) NOT NULL,
    tipe_transaksi VARCHAR(50) NOT NULL,
    tgl_transaksi TIMESTAMPTZ NOT NULL,
    status VARCHAR(50) NOT NULL,
    catatan TEXT,
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

CREATE TRIGGER trg_transaksi_stoks_updated_at 
BEFORE UPDATE ON transaksi_stoks
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
