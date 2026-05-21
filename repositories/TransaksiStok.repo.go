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

func (r *ImplTransaksiStokRepository) GetTransaksiStok(locationId string) ([]dto.GetTransaksiStokDTO, error)

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
					Qty:            item.Qty,
				})
			}

			if err := tx.Create(&transaksiItems).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
