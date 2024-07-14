package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/google/uuid"

	"github.com/otakakot/sample-go-vercel-tips/pkg/openapi"
)

type GetResponse struct {
	Todos []openapi.TodoSchema `json:"todos"`
}

type PostRequest struct {
	Title string `json:"title"`
}

type PostResponse struct {
	Todo openapi.TodoSchema `json:"todo"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		values := r.URL.Query()

		for key, value := range values {
			slog.Info("key: " + key)
			slog.Info("value: " + value[0])
		}

		res := GetResponse{
			Todos: []openapi.TodoSchema{
				{
					Id:        uuid.New(),
					Title:     "Do the dishes",
					Completed: true,
				},
				{
					Id:        uuid.New(),
					Title:     "Walk the dog",
					Completed: false,
				},
			},
		}

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		return
	case http.MethodPost:
		req := PostRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)

			return
		}

		res := PostResponse{
			openapi.TodoSchema{
				Id:        uuid.New(),
				Title:     req.Title,
				Completed: false,
			},
		}

		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		return
	}

	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
