package models

import (
	"time"

	"github.com/google/uuid"
)

type FakturHeader struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`

	NoFaktur      string `gorm:"type:varchar(100);unique;not null" json:"no_faktur"`
	TipeTransaksi string `gorm:"type:varchar(100);not null" json:"tipe_transaksi"`

	TanggalTransaksi time.Time `gorm:"type:date;not null" json:"tgl_transaksi"`

	LokasiID uuid.UUID `gorm:"type:uuid;not null" json:"id_lokasi"`

	IDKontak string `gorm:"type:varchar(100)" json:"id_kontak"`

	Catatan string `gorm:"type:varchar(255)" json:"catatan"`

	Subtotal    int64 `gorm:"type:bigint;default:0" json:"subtotal"`
	TotalDiskon int64 `gorm:"type:bigint;default:0" json:"total_diskon"`
	TotalPajak  int64 `gorm:"type:bigint;default:0" json:"total_pajak"`
	GrandTotal  int64 `gorm:"type:bigint;default:0" json:"grand_total"`

	MetodeBayar string `gorm:"type:varchar(50);not null" json:"metode_bayar"`
	Buktibayar  string `gorm:"type:varchar(100)" json:"bukti_bayar"`

	JumlahBayar int64 `gorm:"type:bigint;default:0" json:"jumlah_bayar"`
	Kembalian   int64 `gorm:"type:bigint;default:0" json:"kembalian"`

	Status string `gorm:"type:varchar(100);not null" json:"status_transaksi"`

	CreatedBy uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
