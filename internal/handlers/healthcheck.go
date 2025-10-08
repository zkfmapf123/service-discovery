package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/zkfmapf123/msa-discovery/internal/dispatcher"
)

func HealthCheck(ctx context.Context, q *dispatcher.Queue, w http.ResponseWriter, r *http.Request) error {

	q.Enqueue(dispatcher.QueueEvents{
		Type:    "logging",
		Data:    map[string]any{"status": "ok"},
		Error:   "",
		Created: time.Now().Unix(),
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))

	return nil
}
