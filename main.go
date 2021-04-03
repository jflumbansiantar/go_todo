package main

import (
	"net/http"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
	fmt
	"io"
)

func init() {
	file, err := OpenFile("todo.log", O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could not open file with error: " + err.Error())
	}

	log.SetOutput(file)
	log.SetFormatter(&log.TextFormatter{})
	log.SetReportCaller(true)
}

func main() {
	log.Info("Starting TODO API")

	//database section
	db := db.DbInit()
	defer db.Close()
	db.DropTableIfExists(&models.TodoItem{}, &models.User{})
	db.AutoMigrate(&models.TodoItem{}, &models.User{})

	router := mux.NewRouter()
	//middleware
	router.Use(middleware.Auth(db))

	router.HandleFunc("/healtcheck", healthCheckHandler).Methods("GET")

	//user handler
	router.HandleFunc("/api/v1/register", handlers.CreateUserHandler(db)).Methods("POST")
	router.HandleFunc("/api/v1/user", handlers.GetListUserHandler(db)).Methods("GET")

	//todo handler
	router.HandleFunc("/api/v1/todo", handlers.CreateTodoHandler(db)).Methods("POST")
	router.HandleFunc("/api/v1/todo", handlers.GetListTodoHandler(db)).Methods("GET")
	router.HandleFunc("/api/v1/todo/{id}", handlers.GetTodoByIDHandler(db)).Methods("GET")
	router.HandleFunc("/api/v1/todo/{id}", handlers.DeleteTodoHandler(db)).Methods("DELETE")
	router.HandleFunc("/api/v1/todo/{id}", handlers.UpdateTodoHandler(db)).Methods("PUT")

	//image handler
	router.HandleFunc("/image/{imageName}", handlers.ShowImageHandler).Methods("GET")
	http.ListenAndServe(":8001", router)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	io.WriteString(w,`{"alive": true}`)
}