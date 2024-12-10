package handlers

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lananolana/webpulse/backend/internal/dto"
)

//go:generate mockery --name=SiteClient --structname=MockSiteClient --output=./ --filename=mock_site_client_test.go --outpkg=handlers
type SiteClient interface {
	GetServiceStats(ctx context.Context, wantedServiceDomain string) dto.ServiceStatsResponse
}

type SiteHandler struct {
	client SiteClient
}

func NewSiteHandler(r *chi.Mux, client SiteClient) *SiteHandler {
	handler := &SiteHandler{
		client: client,
	}

	r.Get("/api/status", handler.GetSiteStatus)

	return handler
}

func (h *SiteHandler) GetSiteStatus(w http.ResponseWriter, r *http.Request) {
	wantedServiceDomain := r.URL.Query().Get("domain")

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
