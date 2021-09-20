package middleware

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (m *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		// ok
		m.Handler.ServeHTTP(w, r)
	} else {
		// error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := &web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
