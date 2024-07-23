package main

import (
	"fmt"
	"net"
)

func main() {
    
    // Escuchar en el puerto 8080
    ln, err := net.Listen("tcp", ":8080")
    if err != nil {
        fmt.Println("Error al escuchar:", err)
        return
    }
    defer ln.Close()

    fmt.Println("Servidor escuchando en el puerto 8080")

    for {
        // Aceptar una conexión
        conn, err := ln.Accept()
        if err != nil {
            fmt.Println("Error al aceptar conexión:", err)
            return
        }

        // Manejar la conexión en una nueva gorutina
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 1024)

    for {
        // Leer los datos del cliente
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("Error al leer del cliente:", err)
            return
        }

        // Imprimir el mensaje recibido
        fmt.Println("Mensaje recibido:", string(buf[:n]))

        // Responder al cliente
        _, err = conn.Write([]byte("Mensaje recibido"))
        if err != nil {
            fmt.Println("Error al escribir al cliente:", err)
            return
        }
    }
}
