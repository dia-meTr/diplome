package httpserver

import (
	"net/http"
	"strconv"
	"strings"

	"oss-backend/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *HTTPServer) getDishByID(w http.ResponseWriter, r *http.Request) {
	dishID, err := uuid.Parse(mux.Vars(r)["dish_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	dish, err := s.dishSrv.GetDishByID(r.Context(), dishID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, dish)
}

func (s *HTTPServer) getDishesByTags(w http.ResponseWriter, r *http.Request) {
	tagStr := r.URL.Query().Get("tag_ids")

	tagIDsStr := strings.Split(tagStr, ",")

	if tagStr == "" {
		dishes, err := s.dishSrv.ListDishes()
		if err != nil {
			s.respondError(w, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, http.StatusOK, dishes)
		return
	}

	tagIDs := make([]int, len(tagIDsStr))

	for i, id := range tagIDsStr {
		tagID, err := strconv.Atoi(id)
		if err != nil {
			s.respondError(w, http.StatusBadRequest, err)
			return
		}

		tagIDs[i] = tagID
	}

	dishes, err := s.dishSrv.GetDishesByTags(r.Context(), tagIDs)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, dishes)
}

func (s *HTTPServer) listDishes(w http.ResponseWriter, r *http.Request) {
	dishes, err := s.dishSrv.ListDishes()
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, dishes)
}

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
