package movie

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/maulibra/simple_movie_api/apimaster/model"
	"github.com/maulibra/simple_movie_api/util"
	"log"
)

type movieRepo struct {
	db *sql.DB
}

func NewMovieRepo(db *sql.DB) IMovieRepository {
	return &movieRepo{db: db}
}

func (m movieRepo) GetMovie() ([]*model.Movie, error) {
	movieList := make([]*model.Movie, 0)
	stmt, err := m.db.Prepare(util.GET_MOVIE)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		movie := model.Movie{}
		err := rows.Scan(&movie.MovieId,&movie.Title,&movie.Duration,&movie.ImageUrl,&movie.Synopsis)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		movieList = append(movieList, &movie)
	}
	return movieList, nil
}

func (m movieRepo) GetMovieByID(id string) (*model.Movie, error) {
	movie := model.Movie{}
	stmt, err := m.db.Prepare(util.GET_MOVIE_BY_ID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&movie.MovieId,&movie.Title,&movie.Duration,&movie.ImageUrl,&movie.Synopsis)
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (m movieRepo) PostMovie(movie *model.Movie) error {
	id := uuid.New()
	movie.MovieId = id.String()
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(util.POST_MOVIE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id, movie.Title,movie.Duration,movie.ImageUrl,movie.Synopsis)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (m movieRepo) UpdateMovie(movie *model.Movie) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(util.UPDATE_MOVIE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(movie.Title,movie.Duration,movie.ImageUrl,movie.Synopsis,movie.MovieId)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (m movieRepo) DeleteMovie(id string) error {
	tx, err := m.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(util.DELETE_MOVIE)
	defer stmt.Close()
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}


