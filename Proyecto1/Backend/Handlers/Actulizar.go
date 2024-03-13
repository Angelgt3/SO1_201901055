package Handlers

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
)

var ChanDatosRAM = make(chan string)
var ChanDatosCPU = make(chan string)

func ActulizarDatosRAM() {
    for {
        archivo, err := os.Open("/proc/ram_so1_1s2024")
        if err != nil {
            fmt.Println("Error al abrir archivo de RAM:", err)
            return
        }
        defer archivo.Close()
        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            fmt.Println("Error al leer archivo de RAM:", err)
            return
        }
        ChanDatosRAM <- string(datos)
        time.Sleep(5 * time.Second)
    }
}

func ActulizarDatosCPU() {
    for {
        archivo, err := os.Open("/proc/cpu_so1_1s2024")
        if err != nil {
            fmt.Println("Error al abrir archivo de CPU:", err)
            return
        }
        defer archivo.Close()
        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            fmt.Println("Error al leer archivo de CPU:", err)
            return
        }
        ChanDatosCPU <- string(datos)
        time.Sleep(5 * time.Second)
    }
}
