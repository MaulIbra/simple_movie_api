package model

type Movie struct {
	MovieId string `json:"movieId"`
	Title string `json:"title"`
	Duration int `json:"duration"`
	ImageUrl string `json:"imageUrl"`
	Synopsis string `json:"synopsis"`
}