package main

import (
	"flag"
	"fmt"
	"net/http"
)

// config stores the configuration of the API
type config struct {
	Addr string // Network port address (e.g., :4000)
	Env  string // Application run environment: development | staging | production
}

// application holds dependencies for the API (configuration, modules, etc.)
type application struct {
	Config config
}

func main() {
	// create a config instance
	cfg := config{}

	// read command line flags used for settings into the config instance
	flag.StringVar(
		&cfg.Addr,
		"addr",
		":4000",
		"Network port address",
	)
	flag.StringVar(
		&cfg.Env,
		"env",
		"development",
		"Running Environment. development|staging|production",
	)
	flag.Parse()

	// create application instance and inject config
	app := application{
		Config: cfg,
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
