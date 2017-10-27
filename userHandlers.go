package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"encoding/base64"

	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

// UserHandler contains http handlers for users
type UserHandler struct {
	userRepo *UserRepository
}

// NewUserHandlers gets a new instance of user handlers
func NewUserHandlers() *UserHandler {
	return &UserHandler{userRepo: NewUserRepo()}
}

// show shows the details of a particular User
func (u *UserHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["userId"])
	if converr != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	user, err := u.userRepo.getUser(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve User")
		return
	}

	RespondWithJSON(w, http.StatusOK, user)
}

// SignupRequestModel is a model used to create new users
type SignupRequestModel struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// signup creates a new user with a given password
func (u *UserHandler) signup(w http.ResponseWriter, r *http.Request) {
	var signup SignupRequestModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&signup); err != nil {
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

	u.userRepo.createUser(&user)
	RespondWithJSON(w, http.StatusCreated, nil)
}

// update accepts a JSON object and updates the matching User
func (u *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	var user User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid User Data")
		return
	}

	u.userRepo.updateUser(&user)
	RespondWithJSON(w, http.StatusCreated, nil)
}
