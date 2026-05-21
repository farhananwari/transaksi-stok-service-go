package models

type StokOpnameItems struct {
	ID                 string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StokOpnameHeaderID string `gorm:"type:uuid;" json:"stok_opname_header_id"`
	BarangID           string `gorm:"type:uuid;" json:"barang_id"`
	StokBalancesID     string `gorm:"type:uuid;" json:"stok_balances_id"`
	StokOpname         int    `gorm:"type:int;" json:"stok_opname"`
	Catatan            string `gorm:"type:varchar(255);not null;" json:"catatan"`
}
