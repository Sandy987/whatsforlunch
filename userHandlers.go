package main

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(vars["userId"])
	if err != nil {
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

// LoginRequestModel is a model used to login
type LoginRequestModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserHandler) login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequestModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&login); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Login")
		return
	}

	hash, err := u.userRepo.getHashForUsername(login.Username)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Login")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(login.Password))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Login")
		return
	}

	token, err := getSignedTokenForUser(login.Username)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Login")
		return
	}

	RespondWithJSON(w, http.StatusOK, token)
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

	if err = u.userRepo.createUser(&user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to signup")
		return
	}

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

	if err := u.userRepo.updateUser(&user); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to update user")
		return
	}

	RespondWithJSON(w, http.StatusCreated, nil)
}
