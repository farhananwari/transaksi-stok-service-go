package models

type TransaksiStokItem struct {
	ID             string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TrasaksiStokID string `gorm:"type:uuid;not null;" json:"transaksi_stok_id"`
	BarangID       string `gorm:"type:uuid;not null;" json:"barang_id"`
	Jumlah         int    `gorm:"type:int;" json:"jumlah"`
}
