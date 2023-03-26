package database

import (
	"bufio"
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	file, err := os.Open("./dbconn.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	var textFile string
	for scanner.Scan() {
		textFile = scanner.Text()
	}

	db, err := sql.Open("mysql", textFile)
	if err != nil {
		panic(err)
	}
	return db
}
