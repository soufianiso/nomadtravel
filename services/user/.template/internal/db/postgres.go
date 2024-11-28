package db

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
    "github.com/soufianiso/nomadtravel/user/configs"
)

func ConnectToPostgres(cfg configs.Config) (*sql.DB, error) {
    connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s",
        cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBAddress)
    
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("failed to open a DB connection: %w", err)
    }

    if err = db.Ping(); err != nil {
        return nil, fmt.Errorf("failed to connect to the database: %w", err)
    }

    return db, nil
}
