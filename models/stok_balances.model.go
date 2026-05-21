package models

import "time"

type StokBalances struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BarangID  string    `gorm:"type:uuid;" json:"barang_id"`
	LokasiID  string    `gorm:"type:uuid;" json:"lokasi_id"`
	Stok      int       `gorm:"type:int;" json:"stok"`
	IsActive  bool      `gorm:"type:bool;" json:"isActive"`
	CreatedBy string    `gorm:"type:uuid;" json:"created_by"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
