package models

import "time"

type StokBalances struct {
	ID         string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BarangID   string    `gorm:"type:uuid;" json:"barang_id"`
	LokasiID   string    `gorm:"type:uuid;" json:"lokasi_id"`
	QtyBalance int       `gorm:"type:int;" json:"qty_balance"`
	QtyOpname  int       `gorm:"type:int;" json:"qty_opname"`
	IsActive   bool      `gorm:"type:bool;" json:"isActive"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
