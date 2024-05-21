package httpserver

import (
	"net/http"

	"oss-backend/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *HTTPServer) getDietByID(w http.ResponseWriter, r *http.Request) {
	dietID, err := uuid.Parse(mux.Vars(r)["diet_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	diet, err := s.dietSrv.GetDietByID(dietID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, diet)
}

func (s *HTTPServer) listDiets(w http.ResponseWriter, r *http.Request) {
	diets, err := s.dietSrv.ListDiets()
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, diets)
}

func (s *HTTPServer) createDiet(w http.ResponseWriter, r *http.Request) {
	var diet models.Diet

	if err := s.decodeJSON(r, &diet); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.dietSrv.CreateDiet(&diet); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, diet)
}

func (s *HTTPServer) updateDiet(w http.ResponseWriter, r *http.Request) {
	var diet models.Diet

	if err := s.decodeJSON(r, &diet); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.dietSrv.UpdateDiet(&diet); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, diet)
}

func (s *HTTPServer) deleteDiet(w http.ResponseWriter, r *http.Request) {
	dietID, err := uuid.Parse(mux.Vars(r)["diet_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.dietSrv.DeleteDiet(dietID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
