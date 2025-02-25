package main

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

func ReadJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxByte := 1_048_578
	body := http.MaxBytesReader(w, r.Body, int64(maxByte))
	decoder := json.NewDecoder(body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(data)
}

func WriteJSONError(w http.ResponseWriter, status int, message string) error {
	type JSONError struct {
		Error string `json:"error"`
	}
	return WriteJSON(w, status, &JSONError{Error: message})
}

func JSONResponse(w http.ResponseWriter, status int, data any) error {
	type JSONEnvlope struct {
		Data any `json:"data"`
	}
	return WriteJSON(w, status, &JSONEnvlope{Data: data})
}
