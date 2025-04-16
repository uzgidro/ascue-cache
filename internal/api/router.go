package api

import (
	"ascue/internal/redisstore"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(store redisstore.Store) http.Handler {
	r := chi.NewRouter()

	r.Get("/api/*", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "*")

		val, err := store.Get(key)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(val))
	})

	return r
}
