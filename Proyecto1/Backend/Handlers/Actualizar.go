package Handlers

import (
    "fmt"
    "io/ioutil"
    "os"
    "time"
    "strings"
    "strconv"
    "Backend/Database"
)

var ChanDatosRAM = make(chan RAMData)
var ChanDatosCPU = make(chan CPUData)

func ActualizarDatosRAM() {
    for {
        archivo, err := os.Open("/proc/ram_so1_1s2024")
        if err != nil {
            fmt.Println("Error al abrir archivo de RAM:", err)
            continue
        }
        defer archivo.Close()
        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            fmt.Println("Error al leer archivo de RAM:", err)
            continue
        }
        ramData := procesarDatosRAM(string(datos))
        ChanDatosRAM <- ramData

        var errDB error
        errDB = Database.InsertDataRAM(ramData.FreeRAMPct, ramData.UsedRAMPct)
        if errDB != nil {
            fmt.Println("Error al actualizar datos de RAM:", errDB)
        }

        time.Sleep(2 * time.Second)
    }
}

func ActualizarDatosCPU() {
    for {
        archivo, err := os.Open("/proc/cpu_so1_1s2024")
        if err != nil {
            fmt.Println("Error al abrir archivo de CPU:", err)
            continue
        }
        defer archivo.Close()
        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            fmt.Println("Error al leer archivo de CPU:", err)
            continue
        }

        cpuData := procesarDatosCPU(string(datos))
        ChanDatosCPU <- cpuData
        
        var errDB error
        errDB = Database.InsertDataCPU(float64(cpuData.FreeCPUPct), float64(cpuData.UsedCPUPct))
        if errDB != nil {
            fmt.Println("Error al actualizar datos de CPU:", errDB)
        }

        time.Sleep(2 * time.Second)
    }
}

func procesarDatosCPU(cpuInfoStr string) CPUData {
    lines := strings.Split(cpuInfoStr, "\n")
    var totalUsage, idleTime int64

    foundCPUInfo := false

    for _, line := range lines {
        if strings.HasPrefix(line, "CPU Info:") {
            foundCPUInfo = true
            continue
        }

        if foundCPUInfo {
            fields := strings.Fields(line)
            if len(fields) < 5 {
                //fmt.Println("Error: Datos de CPU en un formato inesperado:", line)
                break;
            }

            usage, err := strconv.ParseInt(fields[1], 10, 64)
            if err != nil {
                fmt.Println("Error al convertir el valor de uso de la CPU:", err)
                return CPUData{}
            }
            totalUsage += usage

            idle, err := strconv.ParseInt(fields[4], 10, 64)
            if err != nil {
                fmt.Println("Error al convertir el valor de tiempo de inactividad de la CPU:", err)
                return CPUData{}
            }
            idleTime += idle
        }
    }

    TotalCPU := float64(totalUsage + idleTime)
    usagePct := float64(totalUsage) / TotalCPU * 100
    idleTimePct := float64(idleTime) / TotalCPU * 100

    cpuData := CPUData{
        TotalCPU:    int(TotalCPU),
        UsedCPU:     int(totalUsage),
        FreeCPU:     int(idleTime),
        UsedCPUPct:  usagePct,
        FreeCPUPct:  idleTimePct,
    }
    return cpuData
}

func procesarDatosRAM(datos string) RAMData {
    var totalRAM, freeRAM int
    _, err := fmt.Sscanf(datos, "Total RAM: %d\nFree RAM: %d", &totalRAM, &freeRAM)
    if err != nil {
        fmt.Errorf("error al escanear datos de RAM: %v")
        return RAMData{}
    }

    usedRAM := totalRAM - freeRAM
    usedRAMPct := (float64(usedRAM) / float64(totalRAM)) * 100
    freeRAMPct := (float64(freeRAM) / float64(totalRAM)) * 100

    return RAMData{
        TotalRAM:   totalRAM,
        FreeRAM:    freeRAM,
        UsedRAM:    usedRAM,
        UsedRAMPct: usedRAMPct,
        FreeRAMPct: freeRAMPct,
    }
}