package main

import (
	"be/database"
	"be/pkg/mysql"
	"be/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// errEnv := godotenv.Load()
	// if errEnv != nil {
	// 	panic("Failed to load env file")
	// }

	mysql.DatabaseInit()

	database.RunMigration()

	r := mux.NewRouter()

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})

	var port = "8000"
	fmt.Println("server running localhost:8000")

	// Embed the setup allowed in 2 parameter on this below code ...
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
