package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"io"
	"io/ioutil"

	"github.com/gorilla/mux"
)

// LocationHandler contains http handlers
type LocationHandler struct {
	locRepo *LocationRepository
}

// NewLocationHandlers gets a new instance of location handlers
func NewLocationHandlers() *LocationHandler {
	return &LocationHandler{locRepo: NewLocationRepo()}
}

func (h *LocationHandler) list(w http.ResponseWriter, r *http.Request) {
	locations, err := h.locRepo.getAllLocations()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to list locations")
		return
	}
	RespondWithJSON(w, http.StatusOK, locations)
}

// show shows the details of a particular location
func (h *LocationHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["locationId"])
	if converr != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	location, err := h.locRepo.getLocation(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve Location")
		return
	}

	RespondWithJSON(w, http.StatusOK, location)
}

func (h *LocationHandler) create(w http.ResponseWriter, r *http.Request) {
	var location Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Location Data")
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Location Data")
		return
	}

	if err := json.Unmarshal(body, &location); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Location Data")
		return
	}

	h.locRepo.createLocation(&location)
	RespondWithJSON(w, http.StatusCreated, nil)
}

// update accepts a JSON object and updates the matching location
func (h *LocationHandler) update(w http.ResponseWriter, r *http.Request) {
	var location Location
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid location Data")
		return
	}
	if err := r.Body.Close(); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid location Data")
		return
	}

	if err := json.Unmarshal(body, &location); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid location Data")
		return
	}

	h.locRepo.updateLocation(&location)
	RespondWithJSON(w, http.StatusCreated, nil)
}
