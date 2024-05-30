package httpserver

import (
	"fmt"
	"net/http"
	"oss-backend/internal/models"

	"github.com/google/uuid"
)

func (s *HTTPServer) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderID := r.Context().Value("order_id").(uuid.UUID)

	order, err := s.orderSrv.GetOrderByID(r.Context(), orderID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, order)
}

func (s *HTTPServer) getMyOrders(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uuid.UUID)

	orders, err := s.orderSrv.ListUsersOrders(r.Context(), userID)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, orders)

}

func (s *HTTPServer) getOrders(w http.ResponseWriter, r *http.Request) {

	status := r.URL.Query().Get("status")

	deliveryDate := r.URL.Query().Get("delivery_date")

	orders, err := s.orderSrv.ListOrders(r.Context(), status, deliveryDate)
	if err != nil {
		s.respondError(w, http.StatusNotFound, err)
		return
	}

	s.respond(w, http.StatusOK, orders)
}

func (s *HTTPServer) createOrder(w http.ResponseWriter, r *http.Request) {
	var order models.Order

	userID := r.Context().Value("user_id").(uuid.UUID)

	if err := s.decodeJSON(r, &order); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	fmt.Println(order)
	order.UserID = userID

	if err := s.orderSrv.CreateOrder(r.Context(), &order); err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusCreated, order)
}
