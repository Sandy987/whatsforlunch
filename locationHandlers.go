package main

import (
	"encoding/json"
	"net/http"
	"strconv"

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
	id, err := strconv.Atoi(vars["locationId"])
	if err != nil {
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
	var locations []Location
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&locations); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Locations Data")
		return
	}

	for _, loc := range locations {
		if err := h.locRepo.createLocation(&loc); err != nil {
			RespondWithError(w, http.StatusInternalServerError, "Unable to create location")
			return
		}
	}

	RespondWithJSON(w, http.StatusCreated, len(locations))
}

// update accepts a JSON object and updates the matching location
func (h *LocationHandler) update(w http.ResponseWriter, r *http.Request) {
	var location Location
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&location); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Location Data")
		return
	}

	if err := h.locRepo.updateLocation(&location); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to update location")
		return
	}
	RespondWithJSON(w, http.StatusCreated, nil)
}
