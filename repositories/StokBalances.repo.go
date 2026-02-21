package repositories

import (
	"errors"
	"transaksi-stok-service-go/dto"
	"transaksi-stok-service-go/models"

	"gorm.io/gorm"
)

type StokBalancesRepository interface {
	GetBarangWithStok(locationId string) ([]dto.GetBarangStokDTO, error)
	CreateBarangWithStok(data models.StokBalances) (models.StokBalances, error)
	FindByID(id string) (*models.StokBalances, error)
	PatchStokBalance(id string, stokBalance int) error
	IsActiveStok(id string, status bool) error
}

type ImplStokBalancesRepository struct {
	db *gorm.DB
}

func NewStokBalancesRepository(db *gorm.DB) StokBalancesRepository {
	return &ImplStokBalancesRepository{db: db}
}

func (r *ImplStokBalancesRepository) GetBarangWithStok(locationId string) ([]dto.GetBarangStokDTO, error) {
	var result []dto.GetBarangStokDTO

	err := r.db.
		Table("master_barangs mb").
		Select(`
			mb.kode_barang,
			mb.nama,
			sb.qty_system as qty_balances,
			sb.qty_opname
		`).
		Joins("LEFT JOIN stok_balances sb ON sb.barang_id = mb.id").
		Where("sb.lokasi_id = ?", locationId).
		Scan(&result).Error

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *ImplStokBalancesRepository) CreateBarangWithStok(data models.StokBalances) (models.StokBalances, error) {

	err := r.db.Create(&data).Error
	return data, err

}

func (r *ImplStokBalancesRepository) FindByID(id string) (*models.StokBalances, error) {
	var stok models.StokBalances

	err := r.db.
		Table("stok_balances").
		Where("id = ?", id).
		First(&stok).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	return &stok, nil
}

func (r *ImplStokBalancesRepository) PatchStokBalance(id string, stokBalance int) error {
	err := r.db.
		Table("stok_balances").
		Where("id = ?", id).
		Update("qty_balance", stokBalance).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *ImplStokBalancesRepository) IsActiveStok(id string, status bool) error {
	err := r.db.
		Table("stok_balances").
		Where("id = ?", id).
		Update("is_active", status).Error

	if err != nil {
		return err
	}

	return nil
}
