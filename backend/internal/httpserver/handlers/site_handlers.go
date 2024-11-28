package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/internal/httpclient"
)

type SiteClient interface {
	GetServiceStats(ctx context.Context, parsedUrl *url.URL) (dto.ServiceStatsResponse, error)
}

type SiteHandler struct {
	client SiteClient
	mock   bool
}

func NewSiteHandler(r *chi.Mux, client SiteClient, mock bool) *SiteHandler {
	handler := &SiteHandler{
		client: client,
		mock:   mock,
	}

	r.Get("/status", handler.GetSiteStatus)

	return handler
}

func (h *SiteHandler) GetSiteStatus(w http.ResponseWriter, r *http.Request) {
	wantedServiceUrl := r.URL.Query().Get("url")

	if wantedServiceUrl == "" {
		slog.Error("failed to get url param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedUrl, err := url.Parse(wantedServiceUrl)
	if err != nil {
		slog.Error("failed to parse url", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO validate request domain

	if h.mock {
		w.WriteHeader(http.StatusOK)
		if _, err = w.Write([]byte(mockResponseStatus)); err != nil {
			slog.Error("failed to write response", err)
		}

		return
	}

	resp, err := h.client.GetServiceStats(r.Context(), parsedUrl)

	if err != nil && !errors.Is(err, httpclient.ErrTargetUnavailable) {
		slog.Error("failed to get service stats", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(jsonResp); err != nil {
		slog.Error("failed to write response", err)
	}
}
