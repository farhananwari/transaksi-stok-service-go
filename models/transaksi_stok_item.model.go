package models

import "time"

type TransaksiStokItem struct {
	ID             string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	TrasaksiStokID string    `gorm:"type:uuid;" json:"transaksi_stok_id"`
	BarangID       string    `gorm:"type:uuid;" json:"barang_id"`
	Jumlah         int       `gorm:"type:int;" json:"jumlah"`
	Harga          int       `gorm:"type:int;" json:"harga"`
	SubTotal       int       `gorm:"type:int;" json:"sub_total"`
	CreatedBy      string    `gorm:"type:uuid;" json:"created_by"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
}
