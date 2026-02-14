package dto

type GetBarangStokDTO struct {
	KodeBarang string `json:"kode_barang"`
	Nama       string `json:"nama"`
	QtyBalance int    `json:"qty_balances"`
	QtyOpname  int    `json:"qty_opname"`
}

type CreateBarangStokDTO struct {
	KodeBarang string `json:"kode_barang"`
	Nama       string `json:"nama"`
	QtyBalance int    `json:"qty_balances"`
	QtyOpname  int    `json:"qty_opname"`
	LokasiId   string `json:"lokasi_id"`
}
