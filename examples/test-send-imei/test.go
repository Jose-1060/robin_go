package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

func main() {
    // Tu IMEI en formato hexadecimal
    imeiHex := "000F333530333137313739353832343130"

    // Convertir el IMEI hexadecimal a bytes
    imeiBytes, err := hex.DecodeString(imeiHex)
    if err != nil {
        fmt.Println("Error al decodificar el IMEI:", err)
        return
    }

    // La longitud del IMEI en bytes
    imeiLen := len(imeiBytes)

    // Crear un buffer para almacenar la longitud y el IMEI
    buf := make([]byte, 2+imeiLen)

    // Escribir la longitud del IMEI en formato big-endian
    binary.BigEndian.PutUint16(buf[:2], uint16(imeiLen))

    // Copiar los bytes del IMEI al buffer
    copy(buf[2:], imeiBytes)

    // Conectar al servidor
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error al conectar al servidor:", err)
        return
    }
    defer conn.Close()

    // Establecer un tiempo de espera para la escritura
    conn.SetWriteDeadline(time.Now().Add(5 * time.Second))

    // Enviar el buffer al servidor
    _, err = conn.Write(buf)
    if err != nil {
        fmt.Println("Error al enviar el IMEI:", err)
        return
    }

    // Leer la respuesta del servidor
    response := make([]byte, 1)
    conn.SetReadDeadline(time.Now().Add(5 * time.Second))
    _, err = conn.Read(response)
    if err != nil {
        fmt.Println("Error al leer la respuesta:", err)
        return
    }

    // Imprimir la respuesta del servidor
    fmt.Println("Respuesta del servidor:", response)
}
