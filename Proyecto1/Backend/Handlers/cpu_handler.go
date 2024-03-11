package Handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
	"net/http"
)

type ModuleData struct {
	Name  string `json:"name"`
	Data  string `json:"data"`
	Error string `json:"error,omitempty"`
}

func leerDatosCPU(canal chan ModuleData) {
	for {
		archivo, err := os.Open("/proc/cpu_so1_1s2024")
		if err != nil {
			canal <- ModuleData{Name: "CPU", Error: fmt.Sprintf("Error al abrir el archivo del módulo: %v", err)}
			continue
		}
		defer archivo.Close()

		datos, err := ioutil.ReadAll(archivo)
		if err != nil {
			canal <- ModuleData{Name: "CPU", Error: fmt.Sprintf("Error al leer el archivo del módulo: %v", err)}
			continue
		}

		canal <- ModuleData{Name: "CPU", Data: string(datos)}

		time.Sleep(5 * time.Second)
	}
}

func CPUDatosHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	canalCPU := make(chan ModuleData)
	go leerDatosCPU(canalCPU)

	moduleData := <-canalCPU

	jsonData, err := json.Marshal(moduleData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
