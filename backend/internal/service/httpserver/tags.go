package httpserver

import (
	"net/http"
	"strconv"

	"oss-backend/internal/models"

	"github.com/gorilla/mux"
)

func (s *HTTPServer) getTagByID(w http.ResponseWriter, r *http.Request) {
	tagID, err := strconv.Atoi(mux.Vars(r)["tag_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	tag, err := s.tagSrv.GetTagByID(tagID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, tag)
}

func (s *HTTPServer) listTags(w http.ResponseWriter, r *http.Request) {
	tags, err := s.tagSrv.ListTags()
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, tags)
}

func (s *HTTPServer) createTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag

	if err := s.decodeJSON(r, &tag); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.tagSrv.CreateTag(&tag); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, tag)
}

func (s *HTTPServer) updateTag(w http.ResponseWriter, r *http.Request) {
	var tag models.Tag

	if err := s.decodeJSON(r, &tag); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.tagSrv.UpdateTag(&tag); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, tag)
}

func (s *HTTPServer) deleteTag(w http.ResponseWriter, r *http.Request) {
	tagID, err := strconv.Atoi(mux.Vars(r)["tag_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.tagSrv.DeleteTag(tagID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
