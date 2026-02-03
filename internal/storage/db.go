package storage

import (
	"fmt"
	"os"
	"time"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

type DB struct {
	Conn *sql.DB
}

func InitDB() (*DB, error) {
    err := godotenv.Load()
    if err != nil {
    	return nil, fmt.Errorf("error loading .env file")
    }

    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    db, err := sql.Open("pgx", dsn)
    if err != nil {
        return nil, err
    }

    query := `CREATE TABLE IF NOT EXISTS urls (
        id SERIAL PRIMARY KEY,
        original_url TEXT NOT NULL,
	visit_count INTEGER DEFAULT 0,
        short_code TEXT UNIQUE NOT NULL
    );`

    _, err = db.Exec(query)

    db.SetMaxOpenConns(25)
    db.SetMaxIdleConns(25)
    db.SetConnMaxLifetime(5 * time.Minute)

    if err := db.Ping(); err != nil {
	return nil, err
    }
    return &DB{Conn: db}, nil
}
