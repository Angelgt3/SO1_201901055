package main

import (
    "fmt"
    "net/http"
    "Backend/Routes"
    "Backend/Database"
)

func main() {
    Database.ConexionMysql()

    Routes.SetupRoutes()

    fmt.Println("Servidor iniciado. Escuchando en el puerto 8080")
    http.ListenAndServe("0.0.0.0:8080", nil)
}
