package Handlers

import (
	"fmt"
	"net/http"
	"Backend/Database"
)

func RAMDatosHistorico(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	data, err := Database.GetDataRAM()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener datos de la RAM: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
