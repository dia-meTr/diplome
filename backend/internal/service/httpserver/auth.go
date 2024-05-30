package httpserver

import (
	"net/http"
)

func (s *HTTPServer) oauthCallback(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.URL.Query().Get("code")
	if code == "" {
		url := s.authSrv.GetOAuthConsentURL(ctx)
		http.Redirect(w, r, url, http.StatusSeeOther)
		return
	}

	token, err := s.authSrv.ExchangeProvidersCode(ctx, code)
	if err != nil {
		s.respondError(w, http.StatusInternalServerError, err)
		return
	}

	s.respond(w, http.StatusOK, map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
		"token_type":    token.TokenType,
		"expiry":        token.Expiry,
	})
}
