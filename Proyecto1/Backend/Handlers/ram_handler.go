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

func leerDatosRAM(w http.ResponseWriter, r *http.Request) {
    datos, ok := <-ChanDatosRAM
    if !ok {
        http.Error(w, "No hay datos disponibles de RAM", http.StatusNotFound)
        return
    }

    var totalRAM, freeRAM int
    _, err := fmt.Sscanf(datos, "Total RAM: %d\nFree RAM: %d", &totalRAM, &freeRAM)
    if err != nil {
        http.Error(w, "Error al parsear los datos de RAM", http.StatusInternalServerError)
        return
    }

    usedRAM := totalRAM - freeRAM
    usedRAMPct := (float64(usedRAM) / float64(totalRAM)) * 100
    freeRAMPct := (float64(freeRAM) / float64(totalRAM)) * 100

    ramData := RAMData{
        TotalRAM:      totalRAM,
        FreeRAM:       freeRAM,
        UsedRAM:       usedRAM,
        UsedRAMPct:    usedRAMPct,
        FreeRAMPct:    freeRAMPct,
    }

    jsonData, err := json.Marshal(ramData)
    if err != nil {
        http.Error(w, "Error al convertir a JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

func RAMDatosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    leerDatosRAM(w, r)
}
