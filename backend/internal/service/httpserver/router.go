package httpserver

import (
	"net/http"

	"oss-backend/internal/config"
	"oss-backend/internal/models"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

func (s *HTTPServer) newRouter(_ config.Config) *mux.Router {
	var (
		router     = mux.NewRouter()
		api        = router.PathPrefix("/api/v1").Subrouter()
		authorized = api.PathPrefix("/").Subrouter()
		admin      = authorized.PathPrefix("/").Subrouter()
	)

	goth.UseProviders(
		google.New(s.googleOAuthCfg.ClientID, s.googleOAuthCfg.ClientSecret, s.googleOAuthCfg.RedirectURL, s.googleOAuthCfg.Scopes...),
	)

	router.Use(corsMiddleware)
	authorized.Use(s.authMiddleware)

	admin.Use(s.roleMiddleware(models.RoleAdmin))

	api.HandleFunc("/auth/{provider}/callback", s.oauthCallback).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/oauth2callback", s.oauthCallback).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/refresh-token", s.refreshToken).Methods(http.MethodPost, http.MethodOptions)

	api.HandleFunc("/status", s.getStatus).Methods(http.MethodGet, http.MethodOptions)

	api.HandleFunc("/product", s.listProducts).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/product", s.createProduct).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/product", s.updateProduct).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/product/{product_id}", s.deleteProduct).Methods(http.MethodDelete, http.MethodOptions)

	authorized.HandleFunc("/tag", s.listTags).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/tag", s.createTag).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/tag", s.updateTag).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/tag/{tag_id}", s.getTagByID).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/tag/{tag_id}", s.deleteTag).Methods(http.MethodDelete, http.MethodOptions)

	//api.HandleFunc("/dish", s.).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/dish/by_tag", s.getDishesByTags).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/dish", s.getDishesByTags).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/dish", s.createDish).Methods(http.MethodPost, http.MethodOptions)
	api.HandleFunc("/dish", s.updateDish).Methods(http.MethodPut, http.MethodOptions)
	api.HandleFunc("/dish/{dish_id}", s.getDishByID).Methods(http.MethodGet, http.MethodOptions)
	api.HandleFunc("/dish/{dish_id}", s.deleteDish).Methods(http.MethodDelete, http.MethodOptions)
	//api.HandleFunc("/dish-diet", s.AddDietToDish).Methods(http.MethodPost, http.MethodOptions)

	authorized.HandleFunc("/card/clear", s.ClearCard).Methods(http.MethodDelete, http.MethodOptions)
	authorized.HandleFunc("/card/exists", s.ItemExistsById).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/card", s.GetUsersCardByID).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/card", s.AddItemToCard).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/card", s.RemoveItemFromCard).Methods(http.MethodDelete, http.MethodOptions)
	authorized.HandleFunc("/card", s.UpdateItemAmount).Methods(http.MethodPut, http.MethodOptions)

	authorized.HandleFunc("/order", s.getMyOrders).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/order/all", s.getOrders).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/order", s.createOrder).Methods(http.MethodPost, http.MethodOptions)
	authorized.HandleFunc("/order/{order_id}", s.GetOrderByID).Methods(http.MethodGet, http.MethodOptions)
	//authorized.HandleFunc("/order/{order_id}", s.deleteOrder).Methods(http.MethodDelete, http.MethodOptions)
	//authorized.HandleFunc("/order/{order_id}", s.updateOrder).Methods(http.MethodPut, http.MethodOptions)

	authorized.HandleFunc("/profile/me", s.getMe).Methods(http.MethodGet, http.MethodOptions)
	authorized.HandleFunc("/media/upload", s.uploadMedia).Methods(http.MethodPost, http.MethodOptions)

	return router
}

func corsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, req)
	})
}

func (s *HTTPServer) Start() error {
	return http.ListenAndServe(":8080", s.router)
}
