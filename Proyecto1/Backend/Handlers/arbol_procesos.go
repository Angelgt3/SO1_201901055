package Handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

type Proceso struct {
	PID    string    `json:"PID"`
	Name   string    `json:"Name"`
	Hijos  []Proceso `json:"Hijos,omitempty"`
}

func DatosArbol(w http.ResponseWriter, r *http.Request) {
	arbolData := procesarDatosArbol()

	jsonData, err := json.Marshal(arbolData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al serializar datos JSON: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func procesarDatosArbol() Proceso {
	archivo, err := os.Open("/proc/cpu_so1_1s2024")
	if err != nil {
		fmt.Println("Error al abrir archivo de CPU:", err)
	}
	defer archivo.Close()

	datos, err := ioutil.ReadAll(archivo)
	if err != nil {
		fmt.Println("Error al leer archivo de CPU:", err)
	}

	lines := strings.Split(string(datos), "\n")

	var procesos []Proceso
	var procesoActual *Proceso

	for _, line := range lines {
		if strings.HasPrefix(line, "PID:") {
			if procesoActual != nil {
				procesos = append(procesos, *procesoActual)
			}
			fields := strings.Fields(line)
			if len(fields) >= 4 {
				pid := fields[1]
				name := strings.Join(fields[3:], " ")
				procesoActual = &Proceso{PID: pid, Name: name}
			}
		} else if procesoActual != nil && strings.HasPrefix(line, "Children:") {
			for _, childLine := range strings.Split(line, "\n") {
				if strings.HasPrefix(childLine, "PID:") {
					fields := strings.Fields(childLine)
					if len(fields) >= 4 {
						pid := fields[1]
						name := strings.Join(fields[3:], " ")
						procesoActual.Hijos = append(procesoActual.Hijos, Proceso{PID: pid, Name: name})
					}
				}
			}
		}
	}

	// Asegurarse de agregar el último proceso si existe
	if procesoActual != nil {
		procesos = append(procesos, *procesoActual)
	}

	if len(procesos) == 1 {
		return procesos[0]
	}

	return Proceso{PID: "", Name: "root", Hijos: procesos}
}
