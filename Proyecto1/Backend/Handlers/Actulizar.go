package Handlers

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "strings"
    "strconv"
)

var ChanDatosRAM = make(chan string)
var ChanDatosCPU = make(chan CPUData)

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
        
        cpuData := procesarDatosCPU(string(datos))
        ChanDatosCPU <- cpuData
        time.Sleep(5 * time.Second)
    }
}

func procesarDatosCPU(cpuInfoStr string) CPUData {
    lines := strings.Split(cpuInfoStr, "\n")
    cpuInfoFields := strings.Fields(lines[0])[1:]
    var totalUsage, idleTime int64
    for i, value := range cpuInfoFields {
        val, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            fmt.Println("Error al convertir el valor:", err)
            return CPUData{}
        }
        if i == 0 {
            totalUsage += val
        } else if i == 3 {
            idleTime += val
        }
    }
    TotalCPU := float64(totalUsage + idleTime)
    usagePct := float64(totalUsage) / TotalCPU * 100
    idleTimePct := float64(idleTime) / TotalCPU * 100

    return CPUData{
        TotalCPU:    int(TotalCPU),
        UsedCPU:     int(totalUsage),
        FreeCPU:     int(idleTime),
        UsedCPUPct:  usagePct,
        FreeCPUPct:  idleTimePct,
    }
}
