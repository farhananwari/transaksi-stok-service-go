package services

import (
	"transaksi-stok-service-go/dto"
	"transaksi-stok-service-go/models"
	"transaksi-stok-service-go/repositories"
)

type StokBalancesService interface {
	GetBarangWithStok(locationId string) ([]dto.GetBarangStokDTO, error)
	CreateStokBalance(data dto.CreateBarangStokDTO) (dto.CreateBarangStokDTO, error)
	UpdateStokBalance(id string, stokBalance int) error
}

type ImplStokBalancesService struct {
	stokBalancesRepo repositories.StokBalancesRepository
	masterDataRepo   repositories.MasterDataRepository
	transaksiStok    repositories.TransaksiStokRepository
}

func NewStokBalancesService(stokBalancesRepo repositories.StokBalancesRepository, masterDataRepo repositories.MasterDataRepository) StokBalancesService {
	return &ImplStokBalancesService{
		stokBalancesRepo: stokBalancesRepo,
		masterDataRepo:   masterDataRepo,
	}
}

func (s *ImplStokBalancesService) GetBarangWithStok(locationId string) ([]dto.GetBarangStokDTO, error) {
	if _, err := s.masterDataRepo.FindLokasi(locationId); err != nil {
		return []dto.GetBarangStokDTO{}, err
	}
	return s.stokBalancesRepo.GetBarangWithStok(locationId)
}

func (s *ImplStokBalancesService) CreateStokBalance(data dto.CreateBarangStokDTO) (dto.CreateBarangStokDTO, error) {
	// validasi barang dan lokasi

	if _, err := s.masterDataRepo.FindBarang(data.KodeBarang); err != nil {
		return dto.CreateBarangStokDTO{}, err
	}

	if _, err := s.masterDataRepo.FindLokasi(data.LokasiId); err != nil {
		return dto.CreateBarangStokDTO{}, err
	}

	_, err := s.stokBalancesRepo.CreateBarangWithStok(models.StokBalances{

		BarangID: data.KodeBarang,
		LokasiID: data.LokasiId,
		Stok:     data.Stok,
		IsActive: true,
	})

	return data, err
}

func (s *ImplStokBalancesService) UpdateStokBalance(id string, stokBalance int) error {
	if _, err := s.stokBalancesRepo.FindByID(id); err != nil {
		return err
	}
	return s.stokBalancesRepo.UpdateStokBalance(id, stokBalance)
}
