package controllers

import (
	"encoding/json"
	"net/http"
	models "song-library/internal/models"
)

func (c *Controller) GetLibraryData(w http.ResponseWriter, r *http.Request) {
	filters := models.LibraryFilters{
		Group:       r.URL.Query().Get("group"),
		Song:        r.URL.Query().Get("song"),
		ReleaseDate: r.URL.Query().Get("release_date"),
		Limit:       r.URL.Query().Get("limit"),
		Offset:      r.URL.Query().Get("offset"),
	}

	libraryData, err := c.Service.FetchLibData(filters)
	if err != nil {
		http.Error(w, "Library data not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(libraryData); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
