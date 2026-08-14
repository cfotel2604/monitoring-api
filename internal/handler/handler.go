package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct{}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Register(mux *http.ServeMux) {
	mux.HandleFunc("GET /health", h.health)
	mux.HandleFunc("GET /api/v1/metrics/summary", h.summary)
	mux.HandleFunc("GET /api/v1/metrics/timeseries", h.timeseries)
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"status": "ok",
	})
}

func (h *Handler) summary(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"data": map[string]any{
			"hosts":          3,
			"healthy_hosts":  3,
			"cpu_percent":    27.4,
			"memory_percent": 61.8,
		},
	})
}

func (h *Handler) timeseries(w http.ResponseWriter, _ *http.Request) {
	now := time.Now().UTC().Truncate(time.Minute)

	writeJSON(w, http.StatusOK, map[string]any{
		"data": []map[string]any{
			{"time": now.Add(-2 * time.Minute), "value": 22.1},
			{"time": now.Add(-1 * time.Minute), "value": 25.7},
			{"time": now, "value": 27.4},
		},
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
