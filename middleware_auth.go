package main

import (
	"net/http"

	"github.com/anthonypz/rss-aggregator/internal/auth"
	"github.com/anthonypz/rss-aggregator/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, http.StatusForbidden, "Auth error")
			return
		}

		user, err := cfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Couldn't retrieve user")
			return
		}

		handler(w, r, user)
	}
}
