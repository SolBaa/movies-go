package main

import (
	"fmt"
	"net/http"
)

func (app application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK\n")
	fmt.Fprintf(w, "Environment: %s\n", app.config.env)
	fmt.Fprintf(w, "Version: %s\n", version)
}
