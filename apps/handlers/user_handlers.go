package handlers

import (
	"encoding/json"
	"net/http"
	model "github.com/jflumbansiantar/go_todo/apps/models"
	"github.com/jinzhu/gorm"
)

func CreateUserHandler(db *gorm.DB) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		password := r.FormValue("password")

		newUser := &model.User{Name: name, Password: password}
		db.Create(&newUser)
		result := db.Last(&newUser)

		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(result.Value)

	}
	return http.HandlerFunc(fn)
}

func GetListUserHandler(db *gorm.DB) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		users := []model.User{}
		db.Find(&users)
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
	return http.HandlerFunc(fn)
}