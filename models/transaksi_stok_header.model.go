package models

import (
	"time"
)

type TransaksiStokHeader struct {
	ID               string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NoTransaksi      string     `gorm:"type:varchar(100);unique;not null;" json:"no_transaksi"`
	TipeTransaksi    string     `gorm:"type:varchar(100);not null;" json:"tipe_transaksi"` //example : PENJUALAN, PEMBELIAN, MUTASI, RETUR, RUSAK
	TanggalTransaksi time.Time  `gorm:"type:date;not null;" json:"tgl_transaksi"`
	FromLocationID   string     `gorm:"type:uuid;not null;" json:"from_location_id"`
	ToLocationID     string     `gorm:"type:uuid;not null;" json:"to_location_id"`
	Status           string     `gorm:"type:varchar(100);not null;" json:"status_transaksi"`
	Catatan          string     `gorm:"type:varchar(255);not null;" json:"catatan"`
	CreatedBy        string     `gorm:"type:uuid;not null;" json:"created_by"`
	ReceivedBy       string     `gorm:"type:uuid;" json:"received_by"`
	ApprovedBy       string     `gorm:"type:uuid;" json:"approved_by"` // yang approve dari from_location (pemberi barang)
	ApprovedAt       *time.Time `gorm:"type:timestamp;" json:"approved_at"`
	ReceivedAt       *time.Time `gorm:"type:timestamp;" json:"received_at"`
	CreatedAt        time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
}
