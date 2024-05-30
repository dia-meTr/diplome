package auth

import (
	"oss-backend/internal/config"
	"oss-backend/internal/persistence"
	"oss-backend/internal/service"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Service struct {
	cfg config.Config

	repo persistence.Repo

	userSrv service.User

	googleOAuthCfg *oauth2.Config
}

func New(cfg config.Config, userSrv service.User,
	repo persistence.Repo) *Service {
	return &Service{
		cfg:     cfg,
		userSrv: userSrv,

		repo: repo,

		googleOAuthCfg: &oauth2.Config{
			ClientID:     cfg.Oauth.Google.ClientID,
			ClientSecret: cfg.Oauth.Google.ClientSecret,
			RedirectURL:  cfg.Oauth.Google.RedirectURL,
			Scopes:       cfg.Oauth.Google.Scopes,
			Endpoint:     google.Endpoint,
		},
	}
}
