package Handlers

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)

// No necesitas definir ModuleData aquí, ya que está definido en otro lugar

func leerDatosRAM(w http.ResponseWriter, r *http.Request) {
    for {
        archivo, err := os.Open("/proc/ram_so1_1s2024")
        if err != nil {
            handleError(w, fmt.Sprintf("Error al abrir el archivo del módulo: %v", err), http.StatusInternalServerError)
            return
        }
        defer archivo.Close()

        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            handleError(w, fmt.Sprintf("Error al leer el archivo del módulo: %v", err), http.StatusInternalServerError)
            return
        }

        jsonData, err := json.Marshal(ModuleData{Name: "RAM", Data: string(datos)})
        if err != nil {
            handleError(w, fmt.Sprintf("Error al serializar datos JSON: %v", err), http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonData)

        time.Sleep(5 * time.Second)
    }
}

func RAMDatosHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        handleError(w, "Método no permitido", http.StatusMethodNotAllowed)
        return
    }

    leerDatosRAM(w, r)
}

func handleError(w http.ResponseWriter, message string, statusCode int) {
    http.Error(w, message, statusCode)
}
