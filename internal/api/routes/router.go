package routes

import (
	"errors"
	"net/http"
	"path/filepath"
	controllers "song-library/internal/api/controllers"
	services "song-library/internal/services"
)

func CreateRoutes(service services.ServiceInterface) (http.Handler, error) {
	if service == nil {
		return nil, errors.New("service is nil")
	}

	songLib := controllers.NewController(service)

	mux := http.NewServeMux()

	routes := map[string]map[string]http.HandlerFunc{
		"/api/v1/library": {
			http.MethodGet: songLib.GetLibraryData,
		},
		"/api/v1/library/songs": {
			http.MethodGet:    songLib.GetSongText,
			http.MethodPost:   songLib.CreateSong,
			http.MethodPut:    songLib.UpdateTheSong,
			http.MethodDelete: songLib.DeleteTheSong,
		},
	}

	for path, methods := range routes {
		mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			handler, ok := methods[r.Method]
			if !ok {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
			handler(w, r)
		})
	}

	swaggerUIPath := "internal/api/swagger"
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir(swaggerUIPath))))
	mux.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(".", "openapi.yaml"))
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, filepath.Join(swaggerUIPath, "index.html"))
	})

	return mux, nil
}
