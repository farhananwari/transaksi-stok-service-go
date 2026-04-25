package models

import (
	"time"
)

type TransaksiStokHeader struct {
	ID               string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NoTransaksi      string    `gorm:"type:varchar(100);unique;not null;" json:"no_transaksi"`
	TipeTransaksi    string    `gorm:"type:varchar(100);not null;" json:"tipe_transaksi"` //example : PENJUALAN, PEMBELIAN, MUTASI, RETUR, RUSAK
	TanggalTransaksi time.Time `gorm:"type:date;not null;" json:"tgl_transaksi"`
	FromLocationID   string    `gorm:"type:uuid;" json:"from_location_id"`
	ToLocationID     string    `gorm:"type:uuid;" json:"to_location_id"`
	Status           string    `gorm:"type:varchar(100);not null;" json:"status"`
	Catatan          string    `gorm:"type:varchar(255);not null;" json:"catatan"`
	CreatedBy        string    `gorm:"type:uuid;" json:"created_by"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
