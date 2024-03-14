package Handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type RAMData struct {
    TotalRAM      int     `json:"total_ram"`
    FreeRAM       int     `json:"free_ram"`
    UsedRAM       int     `json:"used_ram"`
    UsedRAMPct    float64 `json:"used_ram_pct"`
    FreeRAMPct    float64 `json:"free_ram_pct"`
}

func RAMDatosActual(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    ramInfo := <-ChanDatosRAM

    jsonData, err := json.Marshal(ramInfo)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error al serializar datos JSON: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}
