package httpserver

import (
	"fmt"
	"net/http"

	"oss-backend/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *HTTPServer) createDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish

	if err := s.decodeJSON(r, &dish); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.dishSrv.CreateDish(r.Context(), &dish); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, dish)
}

func (s *HTTPServer) AddDietToDish(w http.ResponseWriter, r *http.Request) {
	var dishDiet models.DishDiet

	if err := s.decodeJSON(r, &dishDiet); err != nil {
		fmt.Println("1")
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.dishSrv.AddDietToDish(r.Context(), &dishDiet); err != nil {
		fmt.Println("2")
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, dishDiet)
}

func (s *HTTPServer) updateDish(w http.ResponseWriter, r *http.Request) {
	var dish models.Dish

	if err := s.decodeJSON(r, &dish); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.dishSrv.UpdateDish(r.Context(), &dish); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, dish)
}

func (s *HTTPServer) deleteDish(w http.ResponseWriter, r *http.Request) {
	dishID, err := uuid.Parse(mux.Vars(r)["dish_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err = s.dishSrv.DeleteDish(r.Context(), dishID); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
