package server

import (
	"elix/utils"
	"fmt"
	"log"
	"net/http"
	// "github.com/charmbracelet/log"
	"github.com/gorilla/mux"
)

func Server(port interface{}) {
	utils.Logger.Info(fmt.Sprintf("Server Listening: http://localhost:%v", port))
	router := mux.NewRouter()
	router.HandleFunc("/get/{key}", Get)
	router.HandleFunc("/new", Post)
	router.HandleFunc("/update", ReValutate)
	router.HandleFunc("/delete", Delete)
	router.HandleFunc("/clear", Clear)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), router)
	if err != nil {
		log.Fatal(err)
	}
}