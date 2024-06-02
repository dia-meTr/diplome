package httpserver

import (
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (s *HTTPServer) GetUsersCardByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	card, err := s.cardSrv.GetUsersCardByID(r.Context(), userID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, card)
}

func (s *HTTPServer) AddItemToCard(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	dishID, err := uuid.Parse(r.URL.Query().Get("dish_id"))
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.cardSrv.AddItemToCard(r.Context(), userID, dishID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, nil)
}

func (s *HTTPServer) ItemExistsByID(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	dishID, err := uuid.Parse(r.URL.Query().Get("dish_id"))
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	amount, err := s.cardSrv.ItemExistsById(r.Context(), userID, dishID)

	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, amount)
}

func (s *HTTPServer) RemoveItemFromCard(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	dishID, err := uuid.Parse(r.URL.Query().Get("dish_id"))
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.cardSrv.RemoveItemFromCard(r.Context(), userID, dishID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, nil)
}

func (s *HTTPServer) UpdateItemAmount(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	dishID, err := uuid.Parse(r.URL.Query().Get("dish_id"))
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	amount, err := strconv.Atoi(r.URL.Query().Get("amount"))
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.cardSrv.UpdateItemAmount(r.Context(), userID, dishID, amount); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, nil)
}

func (s *HTTPServer) ClearCard(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	err := s.cardSrv.ClearCard(r.Context(), userID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, nil)
}
