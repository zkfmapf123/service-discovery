package utils

import (
	"context"
	"net/http"

	"github.com/zkfmapf123/msa-discovery/internal/dispatcher"
)

func DiscoverHandlers(
	ch *dispatcher.Queue,
	hs ...func(
		ctx context.Context,
		ch *dispatcher.Queue,
		w http.ResponseWriter,
		r *http.Request) error) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		for _, h := range hs {
			if err := h(ctx, ch, w, r); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
