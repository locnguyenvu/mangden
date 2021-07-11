package middleware

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ValidJsonBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if contentType := r.Header.Get("Content-Type"); len(contentType) > 0 && contentType != "application/json" {
			http.Error(w, "Wrong content type", http.StatusBadRequest)
			return
		}

		var dumpFormat map[string]interface{}

		if rawBody, err := ioutil.ReadAll(r.Body); err != nil || json.Unmarshal(rawBody, &dumpFormat) != nil {
			http.Error(w, "Invalid json", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
