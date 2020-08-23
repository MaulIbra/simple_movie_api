package main

import (
	"github.com/maulibra/simple_movie_api/apimaster"
	"github.com/maulibra/simple_movie_api/infrastructure"
	"github.com/maulibra/simple_movie_api/infrastructure/config"
)

func main(){
	env := config.NewEnv()
	db := infrastructure.InitDB(env)
	router := infrastructure.MuxRouter()
	apimaster.Init(router,db)
	infrastructure.ListenServe(router)
}
