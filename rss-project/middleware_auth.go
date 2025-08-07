package main

import (
	"fmt"
	"net/http"

	"github.com/Sudhanshu-Taneja/rsagg/internal/auth"
	"github.com/Sudhanshu-Taneja/rsagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			errorResponse(w, 401, fmt.Sprintf("Unauthorized: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			errorResponse(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}

		handler(w, r, user)
	}

}
