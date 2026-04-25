package models

import "time"

type StokOpnameHeader struct {
	ID            string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	NoTransaksi   string    `gorm:"type:varchar(100);unique;not null;" json:"no_transaksi"`
	TanggalOpname time.Time `gorm:"type:date;not null;" json:"tgl_transaksi"`
	Status        string    `gorm:"type:varchar(100);not null;" json:"status"` //example:draft, diterima, tolak
	Catatan       string    `gorm:"type:varchar(255);not null;" json:"catatan"`
	CreatedBy     string    `gorm:"type:uuid;" json:"created_by"`
	IsActive      bool      `gorm:"type:bool;" json:"isActive"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
