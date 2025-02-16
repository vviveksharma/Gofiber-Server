package db

import (
	"fmt"
	"log"
)

// Migrate creates necessary tables if they don't exist
func Migrate() error{
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL
    );`

    _, err := DB.Exec(query)
    if err != nil {
       log.Println("error while creating the database: ", err)
	   return err
    }
    fmt.Println("Migration completed: users table created")
	return nil
}
