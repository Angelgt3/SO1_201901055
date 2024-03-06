package main

import (
	"fmt"
	"time"
)

// Función para simular la lectura de los módulos y enviar los datos a través de un canal
func leerDatosModulo(nombreModulo string, archivoModulo string, canal chan string) {
	for {
		// Simulamos la lectura de los datos del módulo
		// Aquí iría la lógica real para leer los datos del archivo del módulo
		datos := obtenerDatosModulo(nombreModulo, archivoModulo)

		// Enviamos los datos a través del canal
		canal <- datos

		// Esperamos un tiempo antes de volver a leer los datos (por ejemplo, cada 5 segundos)
		time.Sleep(5 * time.Second)
	}
}

// Función para simular la obtención de datos de un módulo (reemplazar con la lógica real de lectura de módulos)
func obtenerDatosModulo(nombreModulo string, archivoModulo string) string {
	// Aquí iría la lógica real para obtener los datos del módulo
	// Por simplicidad, solo devolvemos un mensaje de ejemplo
	return fmt.Sprintf("Datos obtenidos del módulo %s (%s)", nombreModulo, archivoModulo)
}

func main() {
	// Creamos un canal para recibir los datos de los módulos
	canalDatos := make(chan string)

	// Iniciamos una goroutine para leer los datos del módulo de la CPU en segundo plano
	go leerDatosModulo("CPU", "cpu_so1_1s2024.ko", canalDatos)

	// Iniciamos otra goroutine para leer los datos del módulo de RAM en segundo plano
	go leerDatosModulo("RAM", "ram_so1_1s2024.ko", canalDatos)

	// En el bucle principal, leemos continuamente los datos del canal y los mostramos
	for {
		// Leemos los datos del canal
		datos := <-canalDatos

		// Mostramos los datos recibidos
		fmt.Println("Datos recibidos:", datos)

		// Aquí puedes realizar cualquier otra operación con los datos recibidos
	}
}
