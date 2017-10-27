package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DishHandler contains http handlers
type DishHandler struct {
	dishRepo *DishRepository
}

// NewDishHandlers gets a new instance of Dish handlers
func NewDishHandlers() *DishHandler {
	return &DishHandler{dishRepo: NewDishRepo()}
}

func (h *DishHandler) list(w http.ResponseWriter, r *http.Request) {
	dishes, err := h.dishRepo.getAllDishes()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to list Dishes")
		return
	}
	RespondWithJSON(w, http.StatusOK, dishes)
}

// show shows the details of a particular Dish
func (h *DishHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["dishId"])
	if converr != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	dish, err := h.dishRepo.getDish(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve Dish")
		return
	}

	RespondWithJSON(w, http.StatusOK, dish)
}

func (h *DishHandler) create(w http.ResponseWriter, r *http.Request) {
	var dishes []Dish
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dishes); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Dish Data")
		return
	}

	for _, dish := range dishes {
		h.dishRepo.createDish(&dish)
	}

	RespondWithJSON(w, http.StatusCreated, len(dishes))
}

// update accepts a JSON object and updates the matching Dish
func (h *DishHandler) update(w http.ResponseWriter, r *http.Request) {
	var dish Dish
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&dish); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Dish Data")
		return
	}

	h.dishRepo.updateDish(&dish)
	RespondWithJSON(w, http.StatusCreated, nil)
}
