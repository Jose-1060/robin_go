package main

import (
	"log"
	"net/http"

	usecase "github.com/Jose-1060/robin_go/internal/domain/usecases"
	"github.com/Jose-1060/robin_go/internal/infrastructure/db"
	"github.com/Jose-1060/robin_go/internal/interfaces/controllers"
	"github.com/Jose-1060/robin_go/internal/repository"
)

func main() {
    // Conectar a MongoDB
    client, err := db.NewMongoClient("mongodb://localhost:27017")
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("testdb")

    // Crear repositorio, caso de uso y controlador
    deviceRepo := repository.NewDeviceRepositoryMongo(db, "devices")
    deviceUseCase := usecase.NewDeviceUseCase(deviceRepo)
    deviceController := controllers.NewDeviceController(deviceUseCase)

    // Configurar rutas HTTP
    http.HandleFunc("/device", deviceController.GetDeviceByID)
    http.HandleFunc("/devices", deviceController.GetAllDevices)
    http.HandleFunc("/device/create", deviceController.CreateDevice)
    http.HandleFunc("/device/update", deviceController.UpdateDevice)
    http.HandleFunc("/device/delete", deviceController.DeleteDevice)

    // Iniciar el servidor HTTP
    log.Println("Server is running at :8080")
    http.ListenAndServe(":8080", nil)
}
