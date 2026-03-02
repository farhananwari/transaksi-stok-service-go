package dto

import "time"

type GetTransaksiStokDTO struct {
	ID               string    `json:"id"`
	NoTransaksi      string    `json:"no_transaksi"`
	TipeTransaksi    string    `json:"tipe_transaksi"`
	TanggalTransaksi time.Time `json:"tgl_transaksi"`
	FromLocationID   string    `json:"from_location_id"`
	ToLocationID     string    `json:"to_location_id"`
	Status           string    `json:"status"`
	Catatan          string    `json:"catatan"`
	CreatedBy        string    `json:"created_by"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`

	TransaksiItems []GetTransaksiItemDTO `json:"transaksi_items"`
}

type GetTransaksiItemDTO struct {
	ItemID      string `json:"item_id"`
	TransaksiID string `json:"transaksi_id"`
	KodeBarang  string `json:"kode_barang"`
	Qty         int    `json:"qty"`
	Satuan      string `json:"satuan"`
	HargaSatuan int    `json:"harga_satuan"`
	TotalHarga  int    `json:"total_harga"`
}

type CreateTransaksiStokDTO struct {
	NoTransaksi      string    `json:"no_transaksi" binding:"required"`
	TipeTransaksi    string    `json:"tipe_transaksi" binding:"required"`
	TanggalTransaksi time.Time `json:"tgl_transaksi" binding:"required"`
	FromLocationID   string    `json:"from_location_id"`
	ToLocationID     string    `json:"to_location_id"`
	Status           string    `json:"status" binding:"required"`
	Catatan          string    `json:"catatan"`
	CreatedBy        string    `json:"created_by"`
}

type CreateItemTransaksiStokDTO struct {
	TransaksiID string `json:"transaksi_id" binding:"required"`
	KodeBarang  string `json:"kode_barang" binding:"required"`
	Qty         int    `json:"qty" binding:"required"`
	Satuan      string `json:"satuan" binding:"required"`
	HargaSatuan int    `json:"harga_satuan" binding:"required"`
	TotalHarga  int    `json:"total_harga" binding:"required"`
}

type UpdateTransaksiStokDTO struct {
	NoTransaksi      string    `json:"no_transaksi"`
	TipeTransaksi    string    `json:"tipe_transaksi"`
	TanggalTransaksi time.Time `json:"tgl_transaksi"`
	FromLocationID   string    `json:"from_location_id"`
	ToLocationID     string    `json:"to_location_id"`
	Catatan          string    `json:"catatan"`
}

type UpdateItemTransaksiStokDTO struct {
	TransaksiID string `json:"transaksi_id"`
	KodeBarang  string `json:"kode_barang"`
	Qty         int    `json:"qty"`
	Satuan      string `json:"satuan"`
	HargaSatuan int    `json:"harga_satuan"`
	TotalHarga  int    `json:"total_harga"`
}

type PatchStatusTransaksiStokDTO struct {
	Status string `json:"status" binding:"required"`
}
