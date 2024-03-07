package Backend

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/datos/cpu", CPUDatosHandler)
	http.HandleFunc("/datos/ram", RAMDatosHandler)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
