package httpclient

import (
	"context"
	"errors"
	"github.com/lananolana/webpulse/backend/pkg/dnsvalidator"
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
	*http.Client
	EnableMockResponse bool
}

// New creates a new HttpClient with the provided configuration.
func New(cfg config.HTTPClient, enableMockResponse bool) *HttpClient {
	return &HttpClient{
		Client: &http.Client{
			Timeout:   cfg.Timeout,
			Transport: roundtrippers.NewLogging(),
		},
		EnableMockResponse: enableMockResponse,
	}
}

// GetServiceStats sends a request to the specified domain and returns service statistics.
func (c *HttpClient) GetServiceStats(ctx context.Context, domain string) dto.ServiceStatsResponse {

	if domain == "" {
		slog.Error("empty domain name")
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("empty domain name"),
		}
	}

	if !dnsvalidator.DomainIsValid(domain) {
		slog.Error("Invalid domain name format", slog.String("domain", domain))
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("Invalid domain name format"),
		}
	}

	u, err := parseURL(domain)
	if err != nil {
		slog.Error("Failed to parse URL", sl.Err(err))
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("Failed to parse URL"),
		}
	}

	if c.EnableMockResponse {
		return mockServiceStatResponse
	}

	resp, startRequest, err := c.makeRequest(ctx, u)
	if err != nil {
		slog.Error("Failed to make request", sl.Err(err))
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("Failed to make request"),
		}
	}
	defer closeResponseBody(resp.Body)

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", sl.Err(err))
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("Failed to read response body"),
		}
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
		return dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    Ptr("Failed to get server info"),
		}
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
	return nil, time.Time{}, errors.New("failed to perform request with both https and http schemes")
}

// doRequest performs the HTTP request and returns the response.
func (c *HttpClient) doRequest(ctx context.Context, u *url.URL) (*http.Response, time.Time, error) {
	urlStr := u.String()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlStr, nil)
	if err != nil {
		slog.Error("Failed to create request", sl.Err(err))
		return nil, time.Time{}, err
	}

	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Accept", "*/*")

	startRequest := time.Now()
	resp, err := c.Do(req)
	if err != nil {
		slog.Error("Failed to perform request", sl.Err(err))
		return nil, startRequest, err
	}

	return resp, startRequest, nil
}

func closeResponseBody(body io.ReadCloser) {
	if body != nil {
		_, _ = io.Copy(io.Discard, body)
		_ = body.Close()
	}
}
