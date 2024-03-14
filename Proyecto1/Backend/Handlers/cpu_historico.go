package Handlers

import (
	"fmt"
	"net/http"
	"Backend/Database"
)

func CPUDatosHistorico(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	data, err := Database.GetDataCPU()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al obtener datos de la CPU: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
