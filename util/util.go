package util

import (
	"encoding/json"
	"net/http"
)

// ResponseOk When you prepere data to send you can use json like
// normal way to send data, this data will be automatically wrapped
// by data key
func ResponseOk(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]interface{}{
		"data": data,
	}

	json.NewEncoder(w).Encode(body)
}

// ResponseError When you get error in you code you push body of data
// with error code and message of the error
func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
