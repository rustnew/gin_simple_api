package database 


import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "votre_mot_de_passe"
	dbname   = "person_db"
)

var DB *sql.DB


func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	createTable()
}


func createTable() {
	query := `
	CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(50) NOT NULL,
		last_name VARCHAR(50) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INT CHECK (age >= 0 AND age <= 120),
		gender VARCHAR(10) NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}