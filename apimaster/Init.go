package apimaster

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maulibra/simple_movie_api/apimaster/domain/movie"
)

func Init(router *mux.Router, db *sql.DB)  {
	//movie
	movieRepo := movie.NewMovieRepo(db)
	movieUsecase := movie.NewMovieUsecase(movieRepo)
	movieController := movie.NewMovieController(movieUsecase)
	movieController.Movie(router)

}