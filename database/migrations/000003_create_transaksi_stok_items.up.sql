CREATE TABLE "transaksi_stok_items" (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    transaksi_stok_id UUID NOT NULL,
    barang_id UUID NOT NULL,
    qty INTEGER DEFAULT 0,
    harga NUMERIC(12, 2) DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
)