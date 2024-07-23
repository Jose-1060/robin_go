// internal/infrastructure/persistence/mapper/client_mapper.go
package mappers

import (
	"github.com/Jose-1060/robin_go/internal/domain/entities"
	"github.com/Jose-1060/robin_go/internal/domain/models"
)

func ToDeviceModel(client *entities.Device) *models.DeviceModel {
    if client == nil {
        return nil
    }
    return &models.DeviceModel{
        Imei:        client.Imei,
        Timestamp:   client.Timestamp,
        Lat:         client.Lat,
        Lng:         client.Lng,
        Accuracy:    client.Accuracy,
		Speed: 		 client.Speed,
		Ignition: 	 client.Ignition,
		GsmSignal: 	 client.GsmSignal,
		Battery:     client.Battery,
    }
}

func ToDeviceEntity(model *models.DeviceModel) *entities.Device {
    if model == nil {
        return nil
    }
    return &entities.Device{
        Imei:        model.Imei,
        Timestamp:   model.Timestamp,
        Lat:         model.Lat,
        Lng:        model.Lng,
        Accuracy:    model.Accuracy,
		Speed: 		 model.Speed,
		Ignition: 	 model.Ignition,
		GsmSignal: 	 model.GsmSignal,
		Battery:     model.Battery,
    }
}
