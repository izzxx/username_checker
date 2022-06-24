package middleware

import "net/http"

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()

		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Allow-Methods", "GET")
		header.Add("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers")

		next.ServeHTTP(w, r)
	})
}
