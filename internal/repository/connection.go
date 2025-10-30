package repository

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

// DB_Config holds the configuration settings for the database connection pool.
type DB_Config struct {
	DSN            string // Data Source Name
	MaxCon         int    // Maximum number of open connections
	MaxIdleCon     int    // Maximum number of idle connections
	MaxIdleConTime string // Maximum idle connection time
}

// Models is a wrapper struct that holds instances of all the model structs contained within the application.
// model structs holds database operations for a specific table.
type Models struct {
	Groups GroupModel
}

// New creates a new database connection pool and returns it.
func New(cfg DB_Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)

	if err != nil {
		return nil, err
	}

	// Set database connection pool settings
	db.SetMaxOpenConns(cfg.MaxCon)
	db.SetMaxIdleConns(cfg.MaxIdleCon)

	idleDuration, err := time.ParseDuration(cfg.MaxIdleConTime)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxIdleTime(idleDuration)

	// Create a context with a 5 second timeout for the connection test
	ctx, close := context.WithTimeout(context.Background(), 5*time.Second)
	defer close()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// NewModels returns a Models struct containing instances of the model structs.
func NewModels(db *sql.DB) *Models {
	return &Models{
		Groups: GroupModel{db},
	}
}
