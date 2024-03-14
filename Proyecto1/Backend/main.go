package main

import (
    "fmt"
    "net/http"
    "Backend/Routes"
    "Backend/Database"
    "Backend/Handlers"
    "github.com/rs/cors"
)

func main() {
    Database.ConexionMysql()

    Routes.SetupRoutes()

    go Handlers.ActualizarDatosRAM()
    go Handlers.ActualizarDatosCPU()

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"Origin", "Content-Type", "Accept"},
    })

    handler := c.Handler(http.DefaultServeMux)

    fmt.Println("Servidor iniciado. Escuchando en el puerto 8080")
    http.ListenAndServe(":8080", handler)
}
