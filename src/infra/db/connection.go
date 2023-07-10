package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func DbConection() *sql.DB {

	// caminho do arquivo de banco de dados
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//conexão com o banco sqlite
	connection := currentDir + "/src/infra/db/db.db"
	db, err := sql.Open("sqlite3", connection)
	if err != nil {
		log.Fatal(err)
	}

	query := `CREATE TABLE IF NOT EXISTS Products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name VARCHAR,
		Description VARCHAR,
		Price REAL,
		Quantity INTEGER
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
