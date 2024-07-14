package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/otakakot/sample-go-vercel-tips/pkg/openapi"
)

type GetResponse struct {
	Todo openapi.TodoSchema `json:"todo"`
}

type PutRequest struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type PutResponse struct {
	Todo openapi.TodoSchema `json:"todo"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(strings.TrimPrefix(r.URL.Path, "/todos/"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	switch r.Method {
	case http.MethodGet:
		res := GetResponse{
			Todo: openapi.TodoSchema{
				Id:    id,
				Title: "Do the dishes",
			},
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		return
	case http.MethodPut:
		req := PutRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		res := PutResponse{
			Todo: openapi.TodoSchema{
				Id:    id,
				Title: req.Title,
			},
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		return
	case http.MethodDelete:
		w.WriteHeader(http.StatusNoContent)

		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
