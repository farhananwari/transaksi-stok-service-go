package models

import "github.com/google/uuid"

type FakturItems struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	FakturID    string    `gorm:"type:varchar(100);not null" json:"no_faktur"`
	BarangID    uuid.UUID `gorm:"type:uuid;not null" json:"barang_id"`
	Harga       int64     `gorm:"type:bigint;default:0" json:"harga"`
	Qty         int64     `gorm:"type:bigint;default:0" json:"qty"`
	PajakStatus bool      `gorm:"type:bool;default:true" json:"pajak_status"`
}
