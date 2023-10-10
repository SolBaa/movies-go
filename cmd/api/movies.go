package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/SolBaa/movies-go/internal/data"
)

func (app *application) createMoviesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.logger.Println(err)
		http.NotFound(w, r)
		return
	}
	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Year:      1942,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, movie, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
	// fmt.Fprintf(w, "Show the details of a specific movie %d\n", id)
}
