package httpclient

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"time"

	"github.com/lananolana/webpulse/backend/internal/config"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/pkg/http_tools/roundtrippers"
	"github.com/lananolana/webpulse/backend/pkg/logger/sl"
)

// Ptr returns a pointer to the given value.
func Ptr[T any](v T) *T { return &v }

// HttpClient represents an HTTP client.
type HttpClient struct {
	client *http.Client
}

// New creates a new HttpClient with the provided configuration.
func New(cfg config.HTTPClient) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout:   cfg.Timeout,
			Transport: roundtrippers.NewLogging(),
		},
	}
}

// GetServiceStats sends a request to the specified domain and returns service statistics.
func (c *HttpClient) GetServiceStats(ctx context.Context, domain string) dto.ServiceStatsResponse {
	u, err := parseURL(domain)
	if err != nil {
		slog.Error("Failed to parse URL", sl.Err(err))
		return newFailedCreateRequestResponse()
	}

	resp, startRequest, err := c.makeRequest(ctx, u)
	if err != nil {
		return newTargetUnavailableResponse()
	}
	defer closeResponseBody(resp.Body)

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", sl.Err(err))
		return newTargetUnavailableResponse()
	}

	respTime := time.Since(startRequest).Milliseconds()
	transferSpeedKbps := calculateTransferSpeed(len(bytesResp), respTime)
	responseSizeKb := int64(len(bytesResp)) / 1024

	availability := createAvailability(resp)
	performance := dto.Performance{
		ResponseTimeMs:    Ptr(respTime),
		TransferSpeedKbps: Ptr(transferSpeedKbps),
		ResponseSizeKb:    Ptr(responseSizeKb),
		Optimization:      getOptimization(resp),
	}

	sslInfo := getSSLInfo(resp)
	corsInfo := getCORSInfo(resp)

	startDNS := time.Now()
	dnsRecords := getDNSDetails(u.Host)
	dnsResponseTimeMs := time.Since(startDNS).Milliseconds()

	serverInfo, err := getServerInfo(u.Host, resp, dnsResponseTimeMs, dnsRecords)
	if err != nil {
		slog.Error("Failed to get server info", sl.Err(err))
		return newTargetUnavailableResponse()
	}

	return dto.ServiceStatsResponse{
		Status:       dto.Success,
		Availability: &availability,
		Performance:  &performance,
		Security: &dto.Security{
			SSL:  sslInfo,
			CORS: &corsInfo,
		},
		ServerInfo: serverInfo,
	}
}

// makeRequest attempts to make an HTTP request to the given URL.
func (c *HttpClient) makeRequest(ctx context.Context, u *url.URL) (*http.Response, time.Time, error) {
	if u.Scheme != "" {
		return c.doRequest(ctx, u)
	}

	schemes := []string{"https", "http"}
	for _, scheme := range schemes {
		u.Scheme = scheme
		resp, startRequest, err := c.doRequest(ctx, u)
		if err == nil {
			return resp, startRequest, nil
		}
	}
	slog.Error("Failed to perform request with both https and http schemes")
	return nil, time.Time{}, http.ErrHandlerTimeout
}

// doRequest performs the HTTP request and returns the response.
func (c *HttpClient) doRequest(ctx context.Context, u *url.URL) (*http.Response, time.Time, error) {
	urlStr := u.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		slog.Error("Failed to create request", sl.Err(err))
		return nil, time.Time{}, err
	}

	startRequest := time.Now()
	resp, err := c.client.Do(req)
	if err != nil {
		slog.Error("Failed to perform request", sl.Err(err))
		return nil, startRequest, err
	}

	return resp, startRequest, nil
}

func newTargetUnavailableResponse() dto.ServiceStatsResponse {
	return dto.ServiceStatsResponse{
		Status:       dto.Failed,
		Availability: &dto.Availability{IsAvailable: Ptr(false)},
	}
}

func newFailedCreateRequestResponse() dto.ServiceStatsResponse {
	return dto.ServiceStatsResponse{
		Status: dto.Failed,
	}
}

func closeResponseBody(body io.ReadCloser) {
	if body != nil {
		_, _ = io.Copy(io.Discard, body)
		_ = body.Close()
	}
}
