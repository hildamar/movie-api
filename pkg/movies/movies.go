package movies

// Movie represents a single movie
type Movie struct {
	Title string `json:"Title"`
	Year  string `json:"Year"`
}

// MovieSearcher is the interfaces for anything that searches for movies
type MovieSearcher interface {
	SearchMovies(query string, pages string) ([]Movie, error)
	SearchMoviesSorted(query string, page string) ([]Movie, error)
}
