package middlewares

import (
	"WearStoreAPI/db"
	"WearStoreAPI/pkg/auth"
	"log/slog"
	"net/http"
	"os"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		next.ServeHTTP(w, r)
	})
}

func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				slog.Error(err.Error())
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			slog.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := cookie.Value

		claims, err := auth.CheckToken(token, []byte(os.Getenv("JWT_KEY")))
		if err != nil {
			slog.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if useremail, ok := claims["email"].(string); ok {
			var role string
			row := db.DB.QueryRow(`SELECT role FROM permissions_table WHERE email=$1`, useremail)
			if err := row.Scan(&role); err != nil {
				slog.Error(err.Error())
				w.WriteHeader(http.StatusForbidden)
				return
			}
			if role != "admin" {
				w.WriteHeader(http.StatusForbidden)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
