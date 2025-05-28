package api

import (
	"ascue/internal/redisstore"
	"fmt"
	"github.com/go-chi/cors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(store redisstore.Store) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/api/*", func(w http.ResponseWriter, r *http.Request) {
		key := chi.URLParam(r, "*")

		//resources, err := resource.GetResources()
		//if err != nil {
		//	http.Error(w, err.Error(), http.StatusInternalServerError)
		//	return
		//}

		// remove in case using redis
		//data, err := fetch.GetData(resources[key])
		//if err == nil {
		//	fmt.Fprint(w, string(data))
		//} else {
		val, err := store.Get(key)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(val))
		//}
	})

	return r
}
