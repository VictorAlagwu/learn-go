package middlewares

import (
	"errors"
	"net/http"

	"github.com/victoralagwu/learn-go/projects/rest-api/api/auth"
	"github.com/victoralagwu/learn-go/projects/rest-api/api/responses"
)

//SetMiddlewareJSON :
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

//SetMiddlewareAuthentication :
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return 
		}
		next(w, r)
	}
}