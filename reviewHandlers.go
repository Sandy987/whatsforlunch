package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ReviewHandler contains http handlers
type ReviewHandler struct {
	reviewRepo *ReviewRepository
}

// NewReviewHandlers gets a new instance of Review handlers
func NewReviewHandlers() *ReviewHandler {
	return &ReviewHandler{reviewRepo: NewReviewRepo()}
}

func (h *ReviewHandler) list(w http.ResponseWriter, r *http.Request) {
	reviews, err := h.reviewRepo.getAllReviews()
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to list Reviews")
		return
	}
	RespondWithJSON(w, http.StatusOK, reviews)
}

// show shows the details of a particular Review
func (h *ReviewHandler) show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, converr := strconv.Atoi(vars["reviewId"])
	if converr != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	review, err := h.reviewRepo.getReview(id)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Unable to retrieve Review")
		return
	}

	RespondWithJSON(w, http.StatusOK, review)
}

func (h *ReviewHandler) create(w http.ResponseWriter, r *http.Request) {
	var reviews []Review
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reviews); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Reviews Data")
		return
	}

	for _, loc := range reviews {
		h.reviewRepo.createReview(&loc)
	}

	RespondWithJSON(w, http.StatusCreated, len(reviews))
}

// update accepts a JSON object and updates the matching Review
func (h *ReviewHandler) update(w http.ResponseWriter, r *http.Request) {
	var review Review
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&review); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid Review Data")
		return
	}

	h.reviewRepo.updateReview(&review)
	RespondWithJSON(w, http.StatusCreated, nil)
}
