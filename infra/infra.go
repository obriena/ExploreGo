package infra

import (
	"encoding/json"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	b, _ := json.Marshal(data)

	w.Write(b)
}
