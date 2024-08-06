package middleware

import (
	"net/http"
)

func ResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// improve this part
// func ResponseMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		next.ServeHTTP(w, r)

// 		status := r.Context().Value("Status").(string)
// 		message := r.Context().Value("Message").(string)
// 		data := r.Context().Value("Data")
// 		response := models.Response{
// 			Status:  status,
// 			Message: message,
// 			Data:    data,
// 		}

// 		json.NewEncoder(w).Encode(response)
// 	})
// }

// func isAuthenticated(ctx context.Context) (bool, error) {
// 	// mocking data
// 	return true, nil
// }

// func AuthMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		isAuthenticated, err := isAuthenticated(r.Context())
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}

// 		if !isAuthenticated {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), "user", map[string]interface{}{"Username": "test_user"})
// 		r = r.WithContext(ctx)

// 		next.ServeHTTP(w, r)
// 	})
// }
