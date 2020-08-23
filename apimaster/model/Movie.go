package model

type Movie struct {
	MovieId string `json:"movie_id"`
	Title string `json:"title"`
	Duration int `json:"duration"`
	ImageUrl string `json:"image_url"`
	Synopsis string `json:"synopsis"`
}