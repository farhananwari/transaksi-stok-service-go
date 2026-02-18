package repositories

import "gorm.io/gorm"

type MasterDataRepository interface {
	FindBarang(id string) (bool, error)
	FindLokasi(id string) (bool, error)
}

type ImplMasterDataRepository struct {
	db *gorm.DB
}

func NewMasterDataRepository(db *gorm.DB) MasterDataRepository {
	return &ImplMasterDataRepository{db: db}
}

func (r *ImplMasterDataRepository) FindBarang(id string) (bool, error) {
	var count int64
	err := r.db.
		Table("master_barangs").
		Where("id = ?", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *ImplMasterDataRepository) FindLokasi(id string) (bool, error) {
	var count int64
	err := r.db.
		Table("master_lokasis").
		Where("id = ?", id).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
