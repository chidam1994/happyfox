package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error string `json:"error"`
}

func GetBadReqResponse(w http.ResponseWriter, errMsg string) http.ResponseWriter {
	respBody := &Response{Error: errMsg}
	if err := json.NewEncoder(w).Encode(respBody); err != nil {
		w.Write([]byte(err.Error()))
	}
	return w
}
