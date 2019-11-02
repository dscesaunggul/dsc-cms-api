package jsonhttp

import (
	"encoding/json"
	"net/http"
)

// Encode json
func Encode(w http.ResponseWriter, status int, headers map[string]string, data interface{}) error {
	w.WriteHeader(status)
	for k, v := range headers {
		w.Header().Add(k, v)
	}

	body, err := json.Marshal(data)
	if nil != err {
		return err
	}

	w.Write(body)
	return nil
}
