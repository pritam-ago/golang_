package middleware

import (
	"net/http"
)

func RequireRole(role string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value(RoleKey)
			if userRole == nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			if userRole.(string) != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		}
	}
}
