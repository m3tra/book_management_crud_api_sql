package routes

import (
	"book_management_crud_api_sql/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterStaticRoutes = func(r *mux.Router) {
	r.HandleFunc("/hello", controllers.HelloHandler).Methods("GET")
	r.HandleFunc("/form", controllers.FormHandler).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
}
