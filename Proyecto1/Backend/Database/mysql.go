package Database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConexionMysql() (*sql.DB, error) {
	dbUser := "root"
	dbPassword := "201901055"
	dbHost := "database"
	dbPort := "3306"
	dbName := "proyecto1"

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, fmt.Errorf("error al conectar a la base de datos MySQL: %v", err)
	}

	fmt.Println("Conexión con MySQL exitosa")

	return db, nil
}
