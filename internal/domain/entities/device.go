package entities

// Device representa un dispositivo de seguimiento.
type Device struct {
	// ID es el Imei único del dispositivo.
	Imei string
	// Timestamp es la marca de tiempo del último informe del dispositivo.
	Timestamp int64
	// Lat es la latitud del dispositivo.
	Lat float64
	// Lng es la longitud del dispositivo.
	Lng float64
	// Accuracy es la precisión del informe de ubicación.
	Accuracy float64
	// Speed es la velocidad del dispositivo.
	Speed float64
	// Ignition indica si el encendido del dispositivo está activado.
	Ignition bool
	// GsmSignal es la intensidad de la señal GSM.
	GsmSignal int
	// Battery es el nivel de batería del dispositivo.
	Battery int
}
