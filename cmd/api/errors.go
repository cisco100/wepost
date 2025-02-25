package main

import (
	"log"
	"net/http"
)

func (app *Application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error:: %v path:: %v error:: %v \n", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (app *Application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Request  Error:: %v path:: %v error:: %v \n", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *Application) NotExistError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Exist  Error:: %v path:: %v error:: %v \n", r.Method, r.URL.Path, err.Error())

	WriteJSONError(w, http.StatusBadRequest, "resource not exists")
}
