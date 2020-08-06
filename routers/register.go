package routers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoSantos7/twitter/db"
	"github.com/EduardoSantos7/twitter/models"
)

/**
Register is the func to regiser a user
**/
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in data input: "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Mail is required", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password should has at least 6 characters", 400)
		return
	}

	_, find, _ := db.UserExist(t.Email)

	if find == true {
		http.Error(w, "The email is already in use", 400)
	}

	_, status, err := db.CreateUser(t)
	if err != nil {
		http.Error(w, "An error occurs while trying to insert the user", 400)
	}

	if status == false {
		http.Error(w, "The user couldn't be registered", 400)
	}

	w.WriteHeader(http.StatusCreated)
}
