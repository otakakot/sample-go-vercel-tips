package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	api "github.com/otakakot/sample-go-vercel-tips/api/todos"
)

func TestHandler(t *testing.T) {
	t.Parallel()

	t.Run("success_GET_/todos", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodGet, "/todos?key1=value1&key2=value2", nil)

		api.Handler(w, r)

		t.Cleanup(func() {
			r.Body.Close()
		})

		if w.Code != http.StatusOK {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusOK)
		}

		res := api.GetResponse{}

		if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
			t.Fatal(err)
		}

		t.Logf("res: %+v", res)
	})

	t.Run("success_POST_/todos", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		body := bytes.Buffer{}

		req := api.PostRequest{
			Title: "Do the dishes",
		}

		if err := json.NewEncoder(&body).Encode(req); err != nil {
			t.Fatal(err)

			return
		}

		r := httptest.NewRequest(http.MethodPost, "/todos", &body)

		api.Handler(w, r)

		t.Cleanup(func() {
			r.Body.Close()
		})

		if w.Code != http.StatusCreated {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusCreated)
		}

		res := api.PostResponse{}

		if err := json.NewDecoder(w.Body).Decode(&res); err != nil {
			t.Fatal(err)
		}

		if res.Todo.Title != req.Title {
			t.Errorf("got: %v\nwant: %v", res.Todo.Title, req.Title)
		}
	})

	t.Run("fail_PUT_/todos", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodPut, "/todos", nil)

		api.Handler(w, r)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusMethodNotAllowed)
		}
	})

	t.Run("fail_DELETE_/todos", func(t *testing.T) {
		t.Parallel()

		w := httptest.NewRecorder()

		r := httptest.NewRequest(http.MethodDelete, "/todos", nil)

		api.Handler(w, r)

		if w.Code != http.StatusMethodNotAllowed {
			t.Errorf("got: %v\nwant: %v", w.Code, http.StatusMethodNotAllowed)
		}
	})
}
