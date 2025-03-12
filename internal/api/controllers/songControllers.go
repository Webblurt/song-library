package controllers

import (
	"encoding/json"
	"net/http"
	models "song-library/internal/models"
)

func (c *Controller) GetSongText(w http.ResponseWriter, r *http.Request) {
	filters := models.SongFilters{
		SongID: r.URL.Query().Get("song_id"),
		Group:  r.URL.Query().Get("group"),
		Song:   r.URL.Query().Get("song"),
		Page:   r.URL.Query().Get("page"),
		Limit:  r.URL.Query().Get("limit"),
	}

	if filters.SongID == "" && (filters.Group == "" || filters.Song == "") {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	songLyrics, err := c.Service.FetchSongData(filters)
	if err != nil {
		http.Error(w, "Song not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(songLyrics); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func (c *Controller) CreateSong(w http.ResponseWriter, r *http.Request) {
	var req models.NewSongReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.SongCreation(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song created successfully"})
}

func (c *Controller) UpdateTheSong(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateSongReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.Service.SongUpdating(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song updated successfully"})
}

func (c *Controller) DeleteTheSong(w http.ResponseWriter, r *http.Request) {
	req := models.DeleteSongReq{
		ID: r.URL.Query().Get("song_id"),
	}
	if req.ID == "" {
		http.Error(w, "Missing required parameters", http.StatusBadRequest)
		return
	}

	if err := c.Service.SongDeleting(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Song deleted successfully"})
}
