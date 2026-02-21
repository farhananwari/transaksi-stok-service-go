package models

import "time"

type StokOpnameAdjustment struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StokOpnameID string    `gorm:"type:uuid;" json:"stok_opname_id"`
	AdjDate      time.Time `gorm:"type:date;" json:"adj_date"`
	BarangID     string    `gorm:"type:uuid;" json:"barang_id"`
	LokasiID     string    `gorm:"type:uuid;" json:"lokasi_id"`
	QtyBefore    int       `gorm:"type:int;" json:"qty_before"`
	QtyAfter     int       `gorm:"type:int;" json:"qty_after"`
	AdjQty       int       `gorm:"type:int;" json:"adj_qty"`
	Catatan      string    `gorm:"type:varchar(255);not null;" json:"catatan"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
