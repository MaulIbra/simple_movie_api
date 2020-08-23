package movie

import "github.com/maulibra/simple_movie_api/apimaster/model"

type IMovieRepository interface {
	GetMovie()([]*model.Movie,error)
	GetMovieByID(id string) (*model.Movie, error)
	PostMovie(movie *model.Movie) error
	UpdateMovie(movie *model.Movie) error
	DeleteMovie(id string) error
}
