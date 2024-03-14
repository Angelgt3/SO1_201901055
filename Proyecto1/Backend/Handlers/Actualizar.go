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
            return
        }
        defer archivo.Close()
        datos, err := ioutil.ReadAll(archivo)
        if err != nil {
            fmt.Println("Error al leer archivo de RAM:", err)
            return
        }
        ramData := procesarDatosRAM(string(datos))
        ChanDatosRAM <- ramData


        //insertar en la tabla 
        err = Database.InsertDataRAM(ramData.FreeRAMPct, ramData.UsedRAMPct)
        if err != nil {
            fmt.Println("Error al actualizar datos de RAM:", err)
            return
        }

        time.Sleep(5 * time.Second)
    }
}

func ActualizarDatosCPU() {
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
        
        //insertar en la tabla 
        err = Database.InsertDataCPU(float64(cpuData.FreeCPUPct), float64(cpuData.UsedCPUPct))
        if err != nil {
            fmt.Println("Error al actualizar datos de RAM:", err)
            return
        }

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