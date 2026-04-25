package models

import "time"

type LogStok struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BarangID    string    `gorm:"type:uuid;" json:"barang_id"`
	LokasiID    string    `gorm:"type:uuid;" json:"lokasi_id"`
	TransaksiID string    `gorm:"type:uuid;" json:"transaksi_id"`
	QtyIn       int       `gorm:"type:int;" json:"qty_in"`
	QtyOut      int       `gorm:"type:int;" json:"qty_out"`
	CreatedBy   string    `gorm:"type:uuid;" json:"created_by"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
