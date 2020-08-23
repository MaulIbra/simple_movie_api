package infrastructure

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/maulibra/simple_movie_api/util"
	"log"
	"net/http"
)

func MuxRouter() *mux.Router{
	return mux.NewRouter()
}

func ListenServe(router *mux.Router){
	port := util.GetEnv("serverPort","3000")
	host := util.GetEnv("serverHost","localhost")
	log.Print(fmt.Sprintf("%v:%v",host,port))
	err := http.ListenAndServe(fmt.Sprintf("%v:%v",host,port),router)
	if err != nil {
		panic(err)
	}
}
