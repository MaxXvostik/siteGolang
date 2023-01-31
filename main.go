package main

import (
	"log"
	"net/http"

	"example.com/siteGolang/app/controller"
	"example.com/siteGolang/app/server"
	"github.com/julienschmidt/httprouter"
)

func main() {

	err := server.InitDb()
	if err != nil {
		log.Fatal()
	}

	r := httprouter.New()
	routes(r)

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

func routes(r *httprouter.Router) {

	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)
	r.GET("/user", controller.GetUsers)
	r.POST("/user/add", controller.AddUser)
}
