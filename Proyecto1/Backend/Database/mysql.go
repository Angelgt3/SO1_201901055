package Database

import (
	"database/sql"
	"fmt"
	"os"
    "log"
	_ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
	"encoding/json"
)

func ConexionMysql() (*sql.DB, error) {
	/*
	dbUser := "root"
	dbPassword := "201901055"
	dbHost := "database"
	dbPort := "3306"
	dbName := "proyecto1"
	*/
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos MySQL: %v", err)
	}

	fmt.Println("Conexión con MySQL exitosa")

	return db, nil
}

func InsertDataRAM(libre float64, ocupada float64) error {
    db, err := ConexionMysql()
    if err != nil {
        return fmt.Errorf("error al conectar a la base de datos: %v", err)
    }
	defer db.Close()

    query := "INSERT INTO proyecto1.RAM (libre, ocupada) VALUES (?, ?)"

    _, err = db.Exec(query, libre, ocupada)
    if err != nil {
        return fmt.Errorf("error al insertar datos en la tabla RAM: %v", err)
    }

    fmt.Println("Datos de RAM insertados correctamente")
    return nil
}

func InsertDataCPU(libre float64, ocupada float64) error {
    db, err := ConexionMysql()
    if err != nil {
        return fmt.Errorf("error al conectar a la base de datos: %v", err)
    }
	defer db.Close()
	
    query := "INSERT INTO proyecto1.CPU (libre, ocupada) VALUES (?, ?)"

    _, err = db.Exec(query, libre, ocupada)
    if err != nil {
        return fmt.Errorf("error al insertar datos en la tabla CPU: %v", err)
    }

    fmt.Println("Datos de CPU insertados correctamente")
    return nil
}

type CPUData struct {
    ID      int     `json:"id"`
    Libre   float64 `json:"libre"`
    Ocupada float64 `json:"ocupada"`
}

func GetDataCPU() ([]byte, error) {
    db, err := ConexionMysql()
    if err != nil {
        return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    query := "SELECT * FROM proyecto1.CPU ORDER BY id DESC LIMIT 10"
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error al obtener datos de la tabla CPU: %v", err)
    }
    defer rows.Close()

    var cpuDataList []CPUData

    for rows.Next() {
        var data CPUData
        err := rows.Scan(&data.ID, &data.Libre, &data.Ocupada)
        if err != nil {
            return nil, fmt.Errorf("error al escanear fila de resultados: %v", err)
        }
        cpuDataList = append(cpuDataList, data)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error en las filas de resultados: %v", err)
    }

    jsonData, err := json.Marshal(cpuDataList)
    if err != nil {
        return nil, fmt.Errorf("error al convertir datos a JSON: %v", err)
    }

    return jsonData, nil
}

type RAMData struct {
    ID      int     `json:"id"`
    Libre   float64 `json:"libre"`
    Ocupada float64 `json:"ocupada"`
}

func GetDataRAM() ([]byte, error) {
    db, err := ConexionMysql()
    if err != nil {
        return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    query := "SELECT * FROM proyecto1.RAM ORDER BY id DESC LIMIT 10"
    rows, err := db.Query(query)
    if err != nil {
        return nil, fmt.Errorf("error al obtener datos de la tabla RAM: %v", err)
    }
    defer rows.Close()

    var ramDataList []RAMData

    for rows.Next() {
        var data RAMData
        err := rows.Scan(&data.ID, &data.Libre, &data.Ocupada)
        if err != nil {
            return nil, fmt.Errorf("error al escanear fila de resultados: %v", err)
        }
        ramDataList = append(ramDataList, data)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error en las filas de resultados: %v", err)
    }

    jsonData, err := json.Marshal(ramDataList)
    if err != nil {
        return nil, fmt.Errorf("error al convertir datos a JSON: %v", err)
    }

    return jsonData, nil
}