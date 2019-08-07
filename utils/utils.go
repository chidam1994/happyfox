package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
}

func GetBadReqResponse(w http.ResponseWriter, errMsg string) http.ResponseWriter {
	respBody := &Response{Error: errMsg}
	w.WriteHeader(http.StatusBadRequest)
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.Write([]byte(err.Error()))
	}
	return w
}
func GetSuccessReqResponse(w http.ResponseWriter) http.ResponseWriter {
	respBody := &Response{Success: true}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.Write([]byte(err.Error()))
	}
	return w
}
