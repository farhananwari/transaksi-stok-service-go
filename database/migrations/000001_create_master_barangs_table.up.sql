CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE master_barangs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    kode_barang VARCHAR(100) UNIQUE NOT NULL,
    nama VARCHAR(100) NOT NULL,
    harga_beli NUMERIC NOT NULL,
    harga_jual NUMERIC NOT NULL,
    satuan_barang_id VARCHAR(100) NOT NULL,
    jenis_barang_id VARCHAR(100) NOT NULL,

    is_active BOOLEAN NOT NULL DEFAULT true,
    created_by VARCHAR(100) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_master_barangs_nama
    ON master_barangs(nama);

CREATE INDEX idx_master_barangs_is_active
    ON master_barangs(is_active);

CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_master_barangs_updated_at
BEFORE UPDATE ON master_barangs
FOR EACH ROW
EXECUTE FUNCTION set_updated_at();
