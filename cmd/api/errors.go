package main

import (
	"net/http"
)

func (app *Application) InternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Errorw("Internal Server Error", r.Method, "path::", r.URL.Path, "error::", err.Error())

	WriteJSONError(w, http.StatusInternalServerError, "server encountered a problem")
}

func (app *Application) BadRequestError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Errorw("Bad Request Error", r.Method, "path::", r.URL.Path, "error::", err.Error())

	WriteJSONError(w, http.StatusBadRequest, err.Error())
}

func (app *Application) NotExistError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Errorw("Not Exist Error", r.Method, "path::", r.URL.Path, "error::", err.Error())

	WriteJSONError(w, http.StatusBadRequest, "resource not exists")
}

func (app *Application) ConflictError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Errorw("Conflict Error", r.Method, "path::", r.URL.Path, "error::", err.Error())

	WriteJSONError(w, http.StatusConflict, "resource not exists")
}
