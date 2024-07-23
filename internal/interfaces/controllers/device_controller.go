package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jose-1060/robin_go/internal/domain/entities"
	usecase "github.com/Jose-1060/robin_go/internal/domain/usecases"
)

type DeviceController struct{
	DeviceUseCase *usecase.DeviceUseCase
}

func NewDeviceController(uc *usecase.DeviceUseCase) *DeviceController{
	return &DeviceController{
		DeviceUseCase: uc,
	}
}

func (c *DeviceController) GetDeviceByID (w http.ResponseWriter, r *http.Request){
	imei := r.URL.Query().Get("imei")
	device, err := c.DeviceUseCase.GetDeviceByID(imei)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(device)
}

func (c *DeviceController) GetAllDevices (w http.ResponseWriter, r *http.Request){
	devices, err := c.DeviceUseCase.GetAllDevices()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devices)
}

func (c *DeviceController) CreateDevice (w http.ResponseWriter, r *http.Request){
	var device entities.Device
	json.NewDecoder(r.Body).Decode(&device)
	
	err := c.DeviceUseCase.SaveDevice(&device)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(device)
	
}

func (c *DeviceController) UpdateDevice (w http.ResponseWriter, r *http.Request){
	var device entities.Device
	json.NewDecoder(r.Body).Decode(&device)
	err := c.DeviceUseCase.UpdateDevice(&device)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (c *DeviceController) DeleteDevice (w http.ResponseWriter, r *http.Request){
	imei := r.URL.Query().Get("imei")
	err := c.DeviceUseCase.DeleteDevice(imei)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}