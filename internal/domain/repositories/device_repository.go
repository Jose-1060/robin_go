package repositories

import "github.com/Jose-1060/robin_go/internal/domain/entities"

// DeviceRepository define el contrato para un repositorio de dispositivos.
// Este repositorio es responsable de las operaciones CRUD (Create, Read, Update, Delete)
// para los dispositivos.
type DeviceRepository interface {
	// Save guarda un nuevo dispositivo en el repositorio.
	Save(device *entities.Device) error
	// GetAll devuelve una lista de todos los dispositivos almacenados en el repositorio.
	GetAll() ([]*entities.Device, error)
	// GetByID devuelve un dispositivo específico por su ID.
	GetByID(id string) (*entities.Device, error)
	// Update actualiza un dispositivo existente en el repositorio.
	Update(device *entities.Device) error
	// Delete elimina un dispositivo del repositorio.
	Delete(id string) error
}
