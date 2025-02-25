package main

import (
	"net/http"
	"os"
)

func (app *Application) Meta(w http.ResponseWriter, r *http.Request) {
	type Meta struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Version     string `json:"version"`
		Environment string `json:"environment"`
	}
	name := os.Getenv("PROJECT_NAME")
	description := os.Getenv("PROJECT_DESCRIPTION")
	version := os.Getenv("VERSION")
	environment := os.Getenv("ENVIRONMENT")

	metadata := Meta{Name: name, Description: description, Version: version, Environment: environment}
	if err := JSONResponse(w, http.StatusOK, metadata); err != nil {
		app.InternalServerError(w, r, err)

	}

}
