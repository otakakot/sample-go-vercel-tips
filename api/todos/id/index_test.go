package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	api "github.com/otakakot/sample-go-vercel-tips/api/todos/id"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("success_GET_/todos/{id}", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodGet, "/todos/"+uuid.NewString(), nil)

		api.Handler(w, r)

		t.Cleanup(func() {
			r.Body.Close()
		})

		if w.Code != http.StatusOK {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusOK)
		}

		got := api.GetResponse{}

		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			t.Fatal(err)
		}

		if got.Todo.Title != "Do the dishes" {
			t.Errorf("got: %v\nwant: %v", got.Todo.Title, "Do the dishes")
		}
	})

	t.Run("success_PUT_/todos/{id}", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		body := bytes.Buffer{}

		req := api.PutRequest{
			Title: "Walk the dog",
		}

		if err := json.NewEncoder(&body).Encode(req); err != nil {
			t.Fatal(err)
		}

		r := httptest.NewRequest(http.MethodPut, "/todos/"+uuid.NewString(), &body)

		api.Handler(w, r)

		t.Cleanup(func() {
			r.Body.Close()
		})

		if w.Code != http.StatusOK {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusOK)
		}

		res := api.PutResponse{}

		if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
			t.Fatal(err)
		}

		if res.Todo.Title != "Walk the dog" {
			t.Errorf("got: %v\nwant: %v", res.Todo.Title, "Walk the dog")
		}
	})

	t.Run("success_DELETE_/todos/{id}", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodDelete, "/todos/"+uuid.NewString(), nil)

		api.Handler(w, r)

		if w.Code != http.StatusNoContent {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusNoContent)
		}
	})

	t.Run("fail_POST_/todos/{id}", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPost, "/todos/"+uuid.NewString(), nil)

		api.Handler(w, r)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusMethodNotAllowed)
		}
	})
}
