package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/Abdul4code/FairShare/internal"
	"github.com/Abdul4code/FairShare/internal/repository"
	"github.com/joho/godotenv"
)

// config stores the configuration of the API
type config struct {
	Addr string               // Network port address (e.g., :4000)
	Env  string               // Application run environment: development | staging | production
	DB   repository.DB_Config // configurations for the database connection pool
}

// application holds dependencies for the API (configuration, modules, etc.)
type application struct {
	Config config
	Models *repository.Models
}

func main() {
	// load environmental variables
	godotenv.Load()

	// read database settings from environment
	dsn, err := internal.GetString("dsn")
	if err != nil {
		panic(err)
	}

	maxCon, err := internal.GetEnvInt("db_max_cons")
	if err != nil {
		panic(err)
	}

	maxIdleCon, err := internal.GetEnvInt("db_max_idle_cons")
	if err != nil {
		panic(err)
	}

	maxIdleTime, err := internal.GetString("db_max_idle_time")
	if err != nil {
		panic(err)
	}

	// create a config instance with database settings prepopulated
	cfg := config{
		DB: repository.DB_Config{
			DSN:            dsn,
			MaxCon:         maxCon,
			MaxIdleCon:     maxIdleCon,
			MaxIdleConTime: maxIdleTime,
		},
	}

	// read command line flags used for settings into the config instance
	port, _ := internal.GetString("port")
	env, _ := internal.GetString("environment")

	fmt.Println(port)
	fmt.Println(env)

	flag.StringVar(
		&cfg.Addr,
		"addr",
		port,
		"Network port address",
	)
	flag.StringVar(
		&cfg.Env,
		"env",
		env,
		"Running Environment. development|staging|production",
	)
	flag.Parse()

	fmt.Println(cfg)

	// create a database connection
	db, err := repository.New(cfg.DB)
	if err != nil {
		panic(err)
	}

	// create application instance and inject config
	app := application{
		Config: cfg,
		Models: repository.NewModels(db),
	}

	// create a server instance
	server := &http.Server{
		Addr:    app.Config.Addr,
		Handler: app.Router(),
	}

	// Run server
	fmt.Printf(
		"Server running on port %s in %s environment\n",
		app.Config.Addr,
		app.Config.Env,
	)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
