package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/pkg/dnsvalidator"
)

type SiteClient interface {
	GetServiceStats(ctx context.Context, wantedServiceDomain string) dto.ServiceStatsResponse
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

	r.Get("/api/status", handler.GetSiteStatus)

	return handler
}

func (h *SiteHandler) GetSiteStatus(w http.ResponseWriter, r *http.Request) {
	wantedServiceDomain := r.URL.Query().Get("domain")

	if wantedServiceDomain == "" {
		slog.Error("failed to get url param")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(`{"status": "invalid domain name format"}`)); err != nil {
			slog.Error("failed to write response", err)
		}
		return
	}

	if !dnsvalidator.DomainIsValid(wantedServiceDomain) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(`{"status": "invalid domain name"}`)); err != nil {
			slog.Error("failed to write response", err)
		}
		return
	}

	if h.mock {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(mockResponseStatus)); err != nil {
			slog.Error("failed to write response", err)
		}
		return
	}

	resp := h.client.GetServiceStats(r.Context(), wantedServiceDomain)

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal response", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if _, err = w.Write(jsonResp); err != nil {
		slog.Error("failed to write response", err)
	}
}
