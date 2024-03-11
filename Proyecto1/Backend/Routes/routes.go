package Routes

import (
    "net/http"
    "Backend/Handlers"
)

func SetupRoutes() {
    http.HandleFunc("/", Handlers.HelloHandler)
    http.HandleFunc("/ram", Handlers.RAMDatosHandler)
    http.HandleFunc("/cpu", Handlers.CPUDatosHandler)
}
