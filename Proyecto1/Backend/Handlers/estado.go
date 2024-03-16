package Handlers

import (
	"encoding/json"
    "fmt"
    "net/http"
    "os/exec"
    "strconv"
)

type Process struct {
	PID    int      `json:"pid"`
	States []string `json:"states"`
}

var process *Process

func init() {
	process = &Process{}
}

func StartProcess(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("sleep", "infinity")
	err := cmd.Start()
	if err != nil {
		fmt.Print(err)
		http.Error(w, "Error al iniciar el proceso", http.StatusInternalServerError)
		return
	}

	process.PID = cmd.Process.Pid
	process.States = []string{"New"}
	updateProcessState("Ready")
	updateProcessState("Running.")

	fmt.Fprintf(w, "Proceso iniciado con PID: %d y estado en espera", cmd.Process.Pid)
}

func StopProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("kill", "-SIGSTOP", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al detener el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	updateProcessState("Stopped")

	fmt.Fprintf(w, "Proceso con PID %d detenido", pid)
}

func ResumeProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("kill", "-SIGCONT", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al reanudar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	updateProcessState("Running")

	fmt.Fprintf(w, "Proceso con PID %d reanudado", pid)
}

func KillProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	cmd := exec.Command("kill", "-9", strconv.Itoa(pid))
	err = cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error al intentar terminar el proceso con PID %d", pid), http.StatusInternalServerError)
		return
	}

	updateProcessState("Terminated")

	fmt.Fprintf(w, "Proceso con PID %d ha terminado", pid)
}


func updateProcessState(state string) {
	process.States = append(process.States, state)
}

func StatusProcess(w http.ResponseWriter, r *http.Request) {
	pidStr := r.URL.Query().Get("pid")
	if pidStr == "" {
		http.Error(w, "Se requiere el parámetro 'pid'", http.StatusBadRequest)
		return
	}

	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		http.Error(w, "El parámetro 'pid' debe ser un número entero", http.StatusBadRequest)
		return
	}

	for _, process := range processes {
		if process.PID == pid {
			jsonData, err := json.Marshal(process)
			if err != nil {
				http.Error(w, "Error al codificar el proceso en JSON", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)
			return
		}
	}
	http.Error(w, fmt.Sprintf("No se encontró ningún proceso con PID %d", pid), http.StatusNotFound)
}