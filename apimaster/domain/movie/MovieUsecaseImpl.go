package movie

import "github.com/maulibra/simple_movie_api/apimaster/model"

type movieUsecase struct {
	movieRepo IMovieRepository
}

func NewMovieUsecase(repo IMovieRepository) IMovieUsecase  {
	return &movieUsecase{
		movieRepo: repo,
	}
}

func (m movieUsecase) GetMovie() ([]*model.Movie, error) {
	movieList,err := m.movieRepo.GetMovie()
	if err != nil {
		return nil, err
	}
	return movieList,err
}

func (m movieUsecase) GetMovieByID(id string) (*model.Movie, error) {
	movie,err := m.movieRepo.GetMovieByID(id)
	if err != nil {
		return nil, err
	}
	return movie,nil
}

func (m movieUsecase) PostMovie(movie *model.Movie) error {
	err := m.movieRepo.PostMovie(movie)
	if err != nil {
		return err
	}
	return nil
}

func (m movieUsecase) UpdateMovie(movie *model.Movie) error {
	err := m.movieRepo.UpdateMovie(movie)
	if err != nil {
		return err
	}
	return nil
}

func (m movieUsecase) DeleteMovie(id string) error {
	err := m.movieRepo.DeleteMovie(id)
	if err != nil {
		return err
	}
	return nil
}


