package httpserver

import (
	"net/http"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (s *HTTPServer) refreshToken(w http.ResponseWriter, r *http.Request) {
	var reqBody RefreshTokenRequest

	if err := s.decodeJSON(r, &reqBody); err != nil {
		s.respondError(w, http.StatusBadRequest, err)
		return
	}

	token, err := s.authSrv.RefreshToken(r.Context(), reqBody.RefreshToken)
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
