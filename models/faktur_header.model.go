package models

import (
	"time"

	"github.com/google/uuid"
)

type FakturHeader struct {
	ID               uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NoFaktur         string    `gorm:"type:varchar(100);unique;not null;" json:"no_faktur"`
	TipeTransaksi    string    `gorm:"type:varchar(100);not null;" json:"tipe_transaksi"`
	TanggalTransaksi time.Time `gorm:"type:date;not null;" json:"tgl_transaksi"`
	LokasiID         uuid.UUID `gorm:"type:uuid;not null;" json:"id_lokasi"`
	IDMitra          string    `gorm:"type:varchar(100);" json:"id_mitra"`
	Catatan          string    `gorm:"type:varchar(255);not null;" json:"catatan"`
	Status           string    `gorm:"type:varchar(100);not null;" json:"status_transaksi"`
	CreatedBy        string    `gorm:"type:uuid;not null;" json:"created_by"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
