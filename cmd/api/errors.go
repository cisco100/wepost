package main

import (
	"log"
	"net/http"
)

func (app *Application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Internal Server Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (app *Application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Bad Rewuest Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusBadRequest, "bad request")
}

func (app *Application) NotExistError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Not Exists Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusBadRequest, "resource not exists")
}

func (app *Application) ConflictError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Conflict Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusConflict, "resource not exists")
}

func (app *Application) UnauthorizedError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Unauthorized Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)

	WriteJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *Application) UnauthorizedBasicAuthError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Unauthorized-Basic Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)
	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	WriteJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *Application) ForbiddenError(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Forbidden Error,method:: %s, path:: %s, error:: %s\n", r.Method, r.URL.Path, err)
	WriteJSONError(w, http.StatusForbidden, "forbidden")
}
