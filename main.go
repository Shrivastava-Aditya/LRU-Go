package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var cache = NewLRUCache(1024)

func main() {
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)
	http.ListenAndServe(":8080", nil)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "Key parameter is missing", http.StatusBadRequest)
		return
	}

	value, found := cache.Get(key)
	if !found {
		http.Error(w, "Key not found in cache", http.StatusNotFound)
		return
	}

	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"key":   key,
		"value": value,
	})
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	if key == "" || value == "" {
		http.Error(w, "Key or value parameter is missing", http.StatusBadRequest)
		return
	}

	expiration, _ := strconv.Atoi(r.URL.Query().Get("expiration"))
	if expiration <= 0 {
		expiration = 5 // Default expiration of 5 seconds
	}

	cache.Set(key, value, time.Second*time.Duration(expiration))
	jsonResponse(w, http.StatusOK, map[string]interface{}{
		"message": fmt.Sprintf("Key '%s' set successfully with expiration %d seconds", key, expiration),
	})
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers to allow requests from any origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Serve the request
		handler.ServeHTTP(w, r)
	})
}
