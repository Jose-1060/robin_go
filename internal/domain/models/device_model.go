package models

// Device representa un dispositivo de seguimiento.
type DeviceModel struct {
	// ID es el Imei único del dispositivo.
	Imei string `bson:"_id,omitempty"`
	// Timestamp es la marca de tiempo del último informe del dispositivo.
	Timestamp int64 `bson:"timestamp"`
	// Lat es la latitud del dispositivo.
	Lat float64 `bson:"lat"`
	// Lng es la longitud del dispositivo.
	Lng float64 `bson:"lng"`
	// Accuracy es la precisión del informe de ubicación.
	Accuracy float64 `bson:"accuracy"`
	// Speed es la velocidad del dispositivo.
	Speed float64 `bson:"speed"`
	// Ignition indica si el encendido del dispositivo está activado.
	Ignition bool `bson:"ignition"`
	// GsmSignal es la intensidad de la señal GSM.
	GsmSignal int `bson:"gsm_signal"`
	// Battery es el nivel de batería del dispositivo.
	Battery int `bson:"battery"`
}
