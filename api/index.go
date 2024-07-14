package api

import (
	_ "embed"
	"net/http"
)

//go:embed openapi.yaml
var yaml []byte

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Content-Type", "application/yaml")

	w.Write(yaml)
}
