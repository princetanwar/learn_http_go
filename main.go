package main

import (
	"fmt"
	"net/http"
	"time"
)

// HandlerMux combines a ServeMux and a rate limiter middleware
type HandlerMux struct {
	mux         *http.ServeMux
	rateLimiter func(http.ResponseWriter, *http.Request, http.Handler)
}

// ServeHTTP is the main handler method for HandlerMux
func (h *HandlerMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Apply the rate limiter middleware before the mux
	h.rateLimiter(w, r, h.mux)
}

// func SetupRequestMetadata() http.Handler {
// 	// return func(next http.Handler) http.Handler {
// 	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	// 		rmd := defaultRequestMetadata()
// 	// 		ctx := SetRequestMetaData(r.Context(), rmd)
// 	// 		*r = *r.WithContext(ctx)

// 	// 		next.ServeHTTP(w, r)
// 	// 	})
// 	// }
// }

func init() {
	fmt.Println("init")
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "About Page")
	})

	// Dummy rate limiter function
	rateLimiter := func(w http.ResponseWriter, r *http.Request, next http.Handler) {
		// Here, you can implement actual rate limiting logic.
		// For now, just call the next handler.
		fmt.Println("Rate limiting check")
		next.ServeHTTP(w, r)
	}

	// Create a custom HandlerMux
	handlerMux := &HandlerMux{
		mux:         mux,
		rateLimiter: rateLimiter,
	}

	// Configure the server
	server := &http.Server{
		Addr:              ":8080",
		Handler:           handlerMux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	// Start the server
	fmt.Println("Starting server on :8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
