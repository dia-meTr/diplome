package httpserver

import (
	"net/http"

	"oss-backend/internal/models"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func (s *HTTPServer) getProductByID(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(mux.Vars(r)["product_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	product, err := s.productSrv.GetProductByID(productID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, product)
}

func (s *HTTPServer) listProducts(w http.ResponseWriter, r *http.Request) {
	products, err := s.productSrv.ListProducts()
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, products)
}

func (s *HTTPServer) createProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := s.decodeJSON(r, &product); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.productSrv.CreateProduct(&product); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, product)
}

func (s *HTTPServer) updateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := s.decodeJSON(r, &product); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.productSrv.UpdateProduct(&product); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, product)
}

func (s *HTTPServer) deleteProduct(w http.ResponseWriter, r *http.Request) {
	productID, err := uuid.Parse(mux.Vars(r)["product_id"])
	if err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.productSrv.DeleteProduct(productID)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusNoContent, nil)
}
