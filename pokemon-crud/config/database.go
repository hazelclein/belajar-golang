package config

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	
	DB, err = sql.Open("sqlite3", AppConfig.DBPath)
	if err != nil {
		log.Fatal("Error membuka database:", err)
	}

	createTableQuery := `
	CREATE TABLE IF NOT EXISTS pokemon (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		height INTEGER,
		weight INTEGER,
		base_experience INTEGER,
		types TEXT
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal("Error membuat tabel:", err)
	}

	log.Println("Database berhasil diinisialisasi")
}