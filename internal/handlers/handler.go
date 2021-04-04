package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Mishanki/specialist-dz-2/internal/core"
	"github.com/Mishanki/specialist-dz-2/internal/repositories"
	"net/http"
)

var Register = func(w http.ResponseWriter, r *http.Request) {
	//400 и сообщением {"Error" : "User already exists"}
	u := repositories.User{}
	json.NewDecoder(r.Body).Decode(&u)
	w.Header().Set("Content-Type", "application/json")
	valMsg, ok := u.CreateUserValidation()
	if !ok {
		w.WriteHeader(http.StatusConflict)
		msg := core.ErrMsg{Error: valMsg}
		json.NewEncoder(w).Encode(msg)
		return
	}

	id, ok := u.CreateUser()
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		msg := core.ErrMsg{Error: "Error while creating user"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusCreated)
	msg := core.Success{Message: fmt.Sprintf("User was created with id %v. Try to auth.", id)}
	json.NewEncoder(w).Encode(msg)
}

var Auth = func(w http.ResponseWriter, r *http.Request) {
}

var GetAuto = func(w http.ResponseWriter, r *http.Request) {
}

var CreateAuto = func(w http.ResponseWriter, r *http.Request) {
}

var UpdateAuto = func(w http.ResponseWriter, r *http.Request) {
}

var DeleteAuto = func(w http.ResponseWriter, r *http.Request) {
}

var GetStock = func(w http.ResponseWriter, r *http.Request) {
}
