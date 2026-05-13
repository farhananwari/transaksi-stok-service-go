package repositories

import (
	"transaksi-stok-service-go/dto"
	"transaksi-stok-service-go/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransaksiStokRepository interface {
	GetTransaksiStok(locationId string) ([]dto.GetTransaksiStokDTO, error)
	CreateTransaksiStok(dto dto.CreateTransaksiStokDTO, items []dto.CreateItemTransaksiStokDTO) error
}

type ImplTransaksiStokRepository struct {
	db *gorm.DB
}

func NewTransaksiStokRepository(db *gorm.DB) TransaksiStokRepository {
	return &ImplTransaksiStokRepository{db: db}
}

func (r *ImplTransaksiStokRepository) GetTransaksiStok(locationId string) ([]dto.GetTransaksiStokDTO, error) {
	var result []dto.GetTransaksiStokDTO

	err := r.db.Table("transaksi_stoks ts").
		Select(`
			ts.id,
			ts.no_transaksi,
			ts.tipe_transaksi,
			ts.tgl_transaksi,
			ts.from_location_id,
			ts.to_location_id,
			ts.status,
			ts.catatan,
			ts.created_by,
			ts.created_at,
			ts.updated_at,
			ti.id as item_id,
			ti.transaksi_id,
			ti.kode_barang,
			ti.qty,
			ti.satuan,
			ti.harga_satuan,
			ti.total_harga
		`).
		Joins("LEFT JOIN transaksi_items ti ON ti.transaksi_id = ts.id").
		Where("ts.from_location_id = ? OR ts.to_location_id = ?", locationId, locationId).
		Order("ts.created_at DESC").
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	// Group items ke transaksinya masing-masing
	transaksiMap := make(map[string]*dto.GetTransaksiStokDTO)
	var order []string // untuk jaga urutan

	for _, row := range result {
		if _, exists := transaksiMap[row.ID]; !exists {
			transaksiMap[row.ID] = &dto.GetTransaksiStokDTO{
				ID:               row.ID,
				NoTransaksi:      row.NoTransaksi,
				TipeTransaksi:    row.TipeTransaksi,
				TanggalTransaksi: row.TanggalTransaksi,
				FromLocationID:   row.FromLocationID,
				ToLocationID:     row.ToLocationID,
				Catatan:          row.Catatan,
				CreatedBy:        row.CreatedBy,
				CreatedAt:        row.CreatedAt,
				UpdatedAt:        row.UpdatedAt,
				TransaksiItems:   []dto.GetTransaksiItemDTO{},
			}
			order = append(order, row.ID)
		}

	}

	// Susun result sesuai urutan
	for _, id := range order {
		result = append(result, *transaksiMap[id])
	}

	return result, nil
}

func (r *ImplTransaksiStokRepository) CreateTransaksiStok(dto dto.CreateTransaksiStokDTO, items []dto.CreateItemTransaksiStokDTO) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 1. Insert transaksi stok
		transaksi := models.TransaksiStokHeader{
			ID:               uuid.New().String(),
			NoTransaksi:      dto.NoTransaksi,
			TipeTransaksi:    dto.TipeTransaksi,
			TanggalTransaksi: dto.TanggalTransaksi,
			FromLocationID:   dto.FromLocationID,
			ToLocationID:     dto.ToLocationID,
			Catatan:          dto.Catatan,
			CreatedBy:        dto.CreatedBy,
		}

		if err := tx.Create(&transaksi).Error; err != nil {
			return err
		}

		// 2. Insert items kalau ada
		if len(items) > 0 {
			var transaksiItems []models.TransaksiStokItem
			for _, item := range items {
				transaksiItems = append(transaksiItems, models.TransaksiStokItem{
					ID:             uuid.New().String(),
					TrasaksiStokID: transaksi.ID,
					BarangID:       item.KodeBarang,
					Jumlah:         item.Qty,
				})
			}

			if err := tx.Create(&transaksiItems).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
