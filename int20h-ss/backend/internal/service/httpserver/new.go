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

	authSrv       service.Auth
	userSrv       service.User
	subjectSrv    service.Subject
	eventSrv      service.Event
	groupSrv      service.Group
	facultySrv    service.Faculty
	productSrv    service.Product
	mediaSrv      service.Media
	assignmentSrv service.Assignment
	notifierSrv   service.Notifier
	activitySrv   service.Activity
	dietSrv       service.Diet
	tagSrv        service.Tag
	dishSrv       service.Dish
}

func New(cfg config.Config, authSrv service.Auth,
	userSrv service.User, mediaSrv service.Media,
	facultySrv service.Faculty, productSrv service.Product, groupSrv service.Group,
	notifierSrv service.Notifier, subjectSrv service.Subject,
	activitySrv service.Activity, eventSrv service.Event,
	assignmentSrv service.Assignment, dietSrv service.Diet,
	tagSrv service.Tag, dishSrv service.Dish) *HTTPServer {
	server := &HTTPServer{
		cfg:           cfg,
		authSrv:       authSrv,
		userSrv:       userSrv,
		mediaSrv:      mediaSrv,
		facultySrv:    facultySrv,
		productSrv:    productSrv,
		groupSrv:      groupSrv,
		notifierSrv:   notifierSrv,
		activitySrv:   activitySrv,
		subjectSrv:    subjectSrv,
		eventSrv:      eventSrv,
		dietSrv:       dietSrv,
		tagSrv:        tagSrv,
		dishSrv:       dishSrv,
		assignmentSrv: assignmentSrv,
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
