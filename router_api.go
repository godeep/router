package main

import (
	"encoding/json"
	"net/http"
	"runtime"
)

func newApiHandler(rout *Router) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/reload", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		rout.ReloadRoutes()
	})
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		w.Write([]byte("OK"))
	})
	mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		stats := make(map[string]map[string]interface{})
		stats["routes"] = rout.RouteStats()

		json_data, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(json_data)
		w.Write([]byte("\n"))
	})
	mux.HandleFunc("/memory-stats", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		memStats := &runtime.MemStats{}
		runtime.ReadMemStats(memStats)

		json_data, err := json.MarshalIndent(memStats, "", "  ")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(json_data)
		w.Write([]byte("\n"))
	})

	return mux
}
