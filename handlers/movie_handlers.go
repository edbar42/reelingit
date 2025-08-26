package handlers

import (
	"encoding/json"
	"net/http"

	"frontendmasters.com/reelingit/logger"
	"frontendmasters.com/reelingit/models"
)

type MovieHandler struct {
	// TODO: Fill with more data
	logger *logger.Logger
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies := []models.Movie{
		{
			ID:          1,
			TMDB_ID:     181,
			Title:       "The Hacker",
			ReleaseYear: 2022,
			Genres:      []models.Genre{{ID: 1, Name: "Thriller"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 1, FirstName: "Max", LastName: "Deez"}},
		},
		{
			ID:          2,
			TMDB_ID:     182,
			Title:       "Back to the Future",
			ReleaseYear: 1984,
			Genres:      []models.Genre{{ID: 2, Name: "Science Fiction"}},
			Keywords:    []string{},
			Casting:     []models.Actor{{ID: 2, FirstName: "Christopher", LastName: ",Lloyd"}},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		// TODO: Log error
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
