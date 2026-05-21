package dto

type GetBarangStokDTO struct {
	ID             string `json:"id"`
	MasterBarangID string `json:"master_barang_id"`
	KodeBarang     string `json:"kode_barang"`
	Nama           string `json:"nama"`
	Stok           int    `json:"stok"`
}

type CreateBarangStokDTO struct {
	ID             string `json:"id"`
	MasterBarangID string `json:"master_barang_id"`
	KodeBarang     string `json:"kode_barang"`
	Nama           string `json:"nama"`
	Stok           int    `json:"qty_balances"`
	LokasiId       string `json:"lokasi_id"`
}
