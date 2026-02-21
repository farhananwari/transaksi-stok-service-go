package models

type StokOpnameItems struct {
	ID                 string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	StokOpnameHeaderID string `gorm:"type:uuid;" json:"stok_opname_header_id"`
	BarangID           string `gorm:"type:uuid;" json:"barang_id"`
	QtyBalance         int    `gorm:"type:int;" json:"qty_balance"`
	QtyOpname          int    `gorm:"type:int;" json:"qty_opname"`
	QtyDiff            int    `gorm:"type:int;" json:"qty_diff"`
	AdjustmentStatus   string `gorm:"type:varchar(100);not null;" json:"adj_status"`
	Catatan            string `gorm:"type:varchar(255);not null;" json:"catatan"`
}
