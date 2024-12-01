package httpclient

import (
	"context"
	"io"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/lananolana/webpulse/backend/internal/config"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/pkg/http_tools/roundtrippers"
	"github.com/lananolana/webpulse/backend/pkg/logger/sl"
)

func ptrBool(b bool) *bool       { return &b }
func ptrInt64(i int64) *int64    { return &i }
func ptrString(s string) *string { return &s }

// HttpClient represents HTTP client
type HttpClient struct {
	client *http.Client
}

// New creates a new HTTP client
func New(cfg config.HTTPClient) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout:   cfg.Timeout,
			Transport: roundtrippers.NewLogging(),
		},
	}
}

// GetServiceStats sends a request to the specified domain and returns service statistics
func (c *HttpClient) GetServiceStats(ctx context.Context, domain string) dto.ServiceStatsResponse {
	url := "https://" + domain
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		slog.Error("Failed to create request", sl.Err(err))
		return newFailedCreateRequestResponse()
	}

	startRequest := time.Now()
	resp, err := c.client.Do(req)
	if err != nil || resp == nil {
		slog.Error("Failed to perform request", sl.Err(err))
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
		ResponseTimeMs:    ptrInt64(respTime),
		TransferSpeedKbps: ptrInt64(transferSpeedKbps),
		ResponseSizeKb:    ptrInt64(responseSizeKb),
		Optimization:      getOptimization(resp),
	}

	sslInfo := getSSLInfo(resp)
	corsInfo := getCORSInfo(resp)

	startDNS := time.Now()
	dnsRecords := getDNSDetails(domain)
	dnsResponseTimeMs := time.Since(startDNS).Milliseconds()

	serverInfo, err := getServerInfo(domain, resp, dnsResponseTimeMs, dnsRecords)
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

// Helper functions

func newTargetUnavailableResponse() dto.ServiceStatsResponse {
	return dto.ServiceStatsResponse{
		Status:       dto.Failed,
		Availability: &dto.Availability{IsAvailable: ptrBool(false)},
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

func calculateTransferSpeed(bytesReceived int, durationMillis int64) int64 {
	if durationMillis == 0 {
		return 0
	}
	kilobits := float64(bytesReceived*8) / 1000.0
	seconds := float64(durationMillis) / 1000.0
	return int64(kilobits / seconds)
}

func createAvailability(resp *http.Response) dto.Availability {
	return dto.Availability{
		IsAvailable:    ptrBool(true),
		HTTPStatusCode: &resp.StatusCode,
	}
}

func getSSLInfo(resp *http.Response) *dto.SSL {
	if resp.TLS == nil || len(resp.TLS.PeerCertificates) == 0 {
		return nil
	}
	valid := resp.TLS.HandshakeComplete
	expiresAt := resp.TLS.PeerCertificates[0].NotAfter.Unix()
	issuer := resp.TLS.PeerCertificates[0].Issuer.String()

	return &dto.SSL{
		Valid:     &valid,
		ExpiresAt: &expiresAt,
		Issuer:    &issuer,
	}
}

func getCORSInfo(resp *http.Response) dto.CORS {
	allowOrigin := resp.Header.Get("Access-Control-Allow-Origin")
	enabled := allowOrigin != ""
	if !enabled {
		allowOrigin = "*"
	}
	return dto.CORS{
		AllowOrigin: ptrString(allowOrigin),
		Enabled:     ptrBool(enabled),
	}
}

func getDNSDetails(domain string) *dto.DnsRecords {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil
	}

	mxs, err := net.LookupMX(domain)
	if err != nil {
		return nil
	}
	mxDomains := make([]string, len(mxs))
	for i, mx := range mxs {
		mxDomains[i] = mx.Host
	}

	return &dto.DnsRecords{
		A:     ips,
		CNAME: ptrString(domain),
		MX:    mxDomains,
	}
}

func getServerInfo(domain string, resp *http.Response, dnsResponseTimeMs int64, dnsRecords *dto.DnsRecords) (*dto.ServerInfo, error) {
	ipAddresses, err := net.LookupIP(domain)
	if err != nil || len(ipAddresses) == 0 {
		return nil, err
	}

	return &dto.ServerInfo{
		IPAddress:         &ipAddresses[0],
		WebServer:         getServerType(resp),
		DnsResponseTimeMs: &dnsResponseTimeMs,
		DnsRecords:        dnsRecords,
	}, nil
}

func getServerType(resp *http.Response) *string {
	server := resp.Header.Get("Server")
	if server == "" {
		return nil
	}
	return &server
}

func getOptimization(resp *http.Response) *string {
	contentEncoding := resp.Header.Get("Content-Encoding")
	if contentEncoding != "" {
		return &contentEncoding
	}

	none := "none"
	return &none
}
