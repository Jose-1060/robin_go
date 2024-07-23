package usecase

import (
	"github.com/Jose-1060/robin_go/internal/domain/entities"
	"github.com/Jose-1060/robin_go/internal/domain/repositories"
)

type DeviceUseCase struct {
	deviceRepository repositories.DeviceRepository
}

func NewDeviceUseCase(repo repositories.DeviceRepository) *DeviceUseCase {
	return &DeviceUseCase{
		deviceRepository: repo,
	}
}

func (uc *DeviceUseCase) GetDeviceByID(id string) (*entities.Device, error){
	return uc.deviceRepository.GetByID(id)
}

func (uc *DeviceUseCase) GetAllDevices() ([]*entities.Device, error){
	return uc.deviceRepository.GetAll()
}

func (uc *DeviceUseCase) SaveDevice(device *entities.Device) error{
	return uc.deviceRepository.Save(device)
}

func (uc *DeviceUseCase) UpdateDevice(device *entities.Device) error{
	return uc.deviceRepository.Update(device)
}

func (uc *DeviceUseCase) DeleteDevice(id string) error{
	return uc.deviceRepository.Delete(id)
}