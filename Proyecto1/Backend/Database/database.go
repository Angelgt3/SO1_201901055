package Database

import (
    "database/sql"
    "log"
)

var DB *sql.DB

func InitDB() {
	connectionString := "root:201901055@tcp(mysql:3306)/proyecto1"


    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    DB = db
}
