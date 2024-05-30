package httpserver

import (
	"oss-backend/internal/config"
	"oss-backend/internal/service"

	"github.com/gorilla/mux"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type HTTPServer struct {
	cfg            config.Config
	router         *mux.Router
	googleOAuthCfg *oauth2.Config

	authSrv    service.Auth
	userSrv    service.User
	productSrv service.Product
	mediaSrv   service.Media
	tagSrv     service.Tag
	dishSrv    service.Dish
	cardSrv    service.ShoppingCard
	orderSrv   service.Order
}

func New(cfg config.Config, authSrv service.Auth,
	userSrv service.User, mediaSrv service.Media,
	productSrv service.Product, tagSrv service.Tag,
	dishSrv service.Dish, cardSrv service.ShoppingCard,
	orderSrv service.Order) *HTTPServer {
	server := &HTTPServer{
		cfg:        cfg,
		authSrv:    authSrv,
		userSrv:    userSrv,
		mediaSrv:   mediaSrv,
		productSrv: productSrv,
		tagSrv:     tagSrv,
		dishSrv:    dishSrv,
		cardSrv:    cardSrv,
		orderSrv:   orderSrv,
		googleOAuthCfg: &oauth2.Config{
			RedirectURL:  cfg.Oauth.Google.RedirectURL,
			ClientID:     cfg.Oauth.Google.ClientID,
			ClientSecret: cfg.Oauth.Google.ClientSecret,
			Scopes:       cfg.Oauth.Google.Scopes,
			Endpoint:     google.Endpoint,
		},
	}

	server.router = server.newRouter(cfg)

	return server
}
