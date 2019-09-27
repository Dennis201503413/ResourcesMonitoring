package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	common "./common"
)

const (
    STATIC_DIR = "/static/"
    PORT       = "8080"
)

var router = NewRouter()

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
		router.PathPrefix("/css/").Handler(http.StripPrefix("/css/",
		http.FileServer(http.Dir("static/css/"))))
    // Server CSS, JS & Images Statically.
    //router.
    //    PathPrefix(STATIC_DIR).
    //    Handler(http.StripPrefix(STATIC_DIR, http.FileServer(http.Dir("."+STATIC_DIR))))

    return router
}

func main() {

	//cssHandler := http.FileServer(http.Dir("./static/css/"))
	//http.Handle("/static/css/", http.StripPrefix("/static/css/", cssHandler))



	router.HandleFunc("/", common.LoginPageHandler) // GET
	router.HandleFunc("/login", common.LoginHandler).Methods("POST")
	http.Handle("/",router)
	log.Println("Server running on http://localhost:8080")
  log.Fatal(http.ListenAndServe(":8080", nil))

}
