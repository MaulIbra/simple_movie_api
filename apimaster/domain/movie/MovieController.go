package movie

import (
	"github.com/gorilla/mux"
	"github.com/maulibra/simple_movie_api/apimaster/model"
	"github.com/maulibra/simple_movie_api/util"
	"log"
	"net/http"
)

type movieController struct {
	movieUsecase IMovieUsecase
}

func NewMovieController(usecase IMovieUsecase) *movieController  {
	return &movieController{usecase}
}

func (mc *movieController) Movie(r *mux.Router) {
	menu := r.PathPrefix("/movie").Subrouter()
	menu.HandleFunc("", mc.getMovie).Methods(http.MethodGet)
	menu.HandleFunc("/{id}", mc.getMovieById).Methods(http.MethodGet)
	menu.HandleFunc("", mc.createMovie).Methods(http.MethodPost)
	menu.HandleFunc("/{id}", mc.updateMovie).Methods(http.MethodPut)
	menu.HandleFunc("/{id}", mc.deleteMovie).Methods(http.MethodDelete)

}

func (mc *movieController) getMovie(writer http.ResponseWriter, request *http.Request) {
	movieList, err := mc.movieUsecase.GetMovie()
	if err != nil {
		log.Print(err)
		util.HandleResponseError(writer, http.StatusBadGateway)
	}else {
		util.HandleResponseSuccess(writer, http.StatusOK,movieList)
	}
}

func (mc *movieController) getMovieById(writer http.ResponseWriter, request *http.Request) {
	id := util.DecodePathVariabel("id", request)
	movie, err := mc.movieUsecase.GetMovieByID(id)
	if err != nil {
		util.HandleResponseError(writer, http.StatusOK)
	} else {
		util.HandleResponseSuccess(writer, http.StatusOK, movie)
	}
}

func (mc *movieController) createMovie(writer http.ResponseWriter, request *http.Request) {
	var movie model.Movie
	err := util.JsonDecoder(&movie, request)
	if err != nil {
		util.HandleResponseError(writer, http.StatusBadRequest)
	} else {
		err = mc.movieUsecase.PostMovie(&movie)
		if err != nil {
			util.HandleResponseError(writer, http.StatusBadGateway)
		} else {
			util.HandleResponseSuccess(writer, http.StatusCreated,movie)
		}
	}
}

func (mc *movieController) updateMovie(writer http.ResponseWriter, request *http.Request) {
	var movie model.Movie
	id := util.DecodePathVariabel("id", request)
	err := util.JsonDecoder(&movie, request)
	if err != nil {
		log.Print(err)
	}
	movie.MovieId = id
	err = mc.movieUsecase.UpdateMovie(&movie)
	if err != nil {
		util.HandleResponseError(writer, http.StatusBadGateway)
	} else {
		util.HandleResponseSuccess(writer, http.StatusOK, movie)
	}
}

func (mc *movieController) deleteMovie(writer http.ResponseWriter, request *http.Request) {
	id := util.DecodePathVariabel("id", request)
	if len(id) > 0 {
		err := mc.movieUsecase.DeleteMovie(id)
		if err != nil {
			util.HandleResponseError(writer, http.StatusNotFound)
		} else {
			util.HandleResponseError(writer, http.StatusOK)
		}
	} else {
		util.HandleResponseError(writer, http.StatusBadRequest)
	}
}
