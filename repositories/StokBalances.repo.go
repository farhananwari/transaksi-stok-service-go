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
	UpdateStokBalance(id string, stokBalance int) error
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
		Table("stok_balances sb").
		Select(`
			sb.id,
			mb.id as master_barang_id,
			mb.kode_barang,
			mb.nama,
			COALESCE(sb.stok, 0) as stok
		`).
		Joins(`
			JOIN master_barang mb 
			ON mb.id = sb.master_barang_id
		`).
		Where("sb.location_id = ?", locationId).
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

func (r *ImplStokBalancesRepository) UpdateStokBalance(id string, stokBalance int) error {
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
