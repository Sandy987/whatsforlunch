package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"encoding/base64"
	"io"
	"io/ioutil"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

// Show shows the details of a particular User
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["userId"])
	if converr != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	user, err := GetUser(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve User")
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

// Signup creates a new user with a given password
func Signup(w http.ResponseWriter, r *http.Request) {
	var signup SignupRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Signup Data")
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Signup Data")
		return
	}

	if err := json.Unmarshal(body, &signup); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Signup Data")
		return
	}

	user := HashedUser{
		Username:  signup.Username,
		FirstName: signup.FirstName,
		LastName:  signup.LastName,
		Email:     signup.Email,
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(signup.Password), bcrypt.DefaultCost)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to signup")
		return
	}

	user.PasswordHash = base64.StdEncoding.EncodeToString(hash)

	CreateUser(&user)
	RespondWithJSON(w, http.StatusCreated, nil)
}

// Update accepts a JSON object and updates the matching User
func Update(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User Data")
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User Data")
		return
	}

	if err := json.Unmarshal(body, &user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User Data")
		return
	}

	UpdateUser(&user)
	RespondWithJSON(w, http.StatusCreated, nil)
}
