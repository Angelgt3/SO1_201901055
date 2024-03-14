package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Esta estructura representa los datos de la CPU.
type CPUData struct {
	TotalCPU    int     `json:"total_cpu"`
	UsedCPU     int     `json:"used_cpu"`
	FreeCPU     int     `json:"free_cpu"`
	UsedCPUPct  float64 `json:"used_cpu_pct"`
	FreeCPUPct  float64 `json:"free_cpu_pct"`
}

func CPUDatosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	cpuInfo := <-ChanDatosCPU

	jsonData, err := json.Marshal(cpuInfo)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
