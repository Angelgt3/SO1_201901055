package Routes

import (
    "net/http"
    "Backend/Handlers"
)

func SetupRoutes() {
    http.HandleFunc("/", Handlers.HelloHandler)
    http.HandleFunc("/ram", Handlers.RAMDatosActual)
    http.HandleFunc("/cpu", Handlers.CPUDatosActual)
    http.HandleFunc("/historico/cpu", Handlers.CPUDatosHistorico)
    http.HandleFunc("/historico/ram", Handlers.RAMDatosHistorico)
    http.HandleFunc("/arbol", Handlers.DatosArbol)
}
