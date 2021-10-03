package handlers

import (
	"net/http"

	"flyingspheres.com/test/infra"
)

func Health(w http.ResponseWriter, r *http.Request) {
	m := make(map[string]string)
	m["message"] = "Everything seems okay :-)"
	infra.RespondJSON(w, r, m)
}
