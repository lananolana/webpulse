package httpclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/pkg/http_tools/roundtrippers"
)

var (
	ErrTargetUnavailable = errors.New("target unavailable")
)

// HttpClient represents http client
type HttpClient struct {
	client *http.Client
}

// New creates new http client
func New() *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout:   10 * time.Second,
			Transport: roundtrippers.NewLogging(),
		},
	}
}

// GetServiceStats sends request to url and returns response
func (c *HttpClient) GetServiceStats(ctx context.Context, parsedUrl *url.URL) (dto.ServiceStatsResponse, error) {
	// Creating request object
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, parsedUrl.String(), nil)
	if err != nil {
		fmt.Println(err)
		return dto.ServiceStatsResponse{}, err
	}

	start := time.Now()

	// Making request
	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Println(err)
		return wrapTargetUnavailableResponse()
	}

	availability := dto.Availability{
		IsAvailable:    true,
		HTTPStatusCode: &resp.StatusCode,
	}

	performance := dto.Performance{
		ResponseTime:  time.Since(start).String(),
		DownloadSpeed: formatSpeed(int64(float64(resp.ContentLength) / time.Since(start).Seconds())),
	}

	sslInfo := dto.SSL{
		IsValid:              resp.TLS != nil && resp.TLS.HandshakeComplete,
		ExpiresAt:            resp.TLS.PeerCertificates[0].NotAfter,
		Issuer:               resp.TLS.PeerCertificates[0].Issuer.String(),
		CertificateAuthority: resp.TLS.PeerCertificates[0].Subject.String(),
	}

	dns := getDNSDetails(parsedUrl.Hostname())

	// Deferred response body discard and close
	defer func(Body io.ReadCloser) {
		_, _ = io.Copy(io.Discard, Body)
		_ = Body.Close()
	}(resp.Body)

	return dto.ServiceStatsResponse{
		Target:       resp.Request.URL.String(),
		Availability: availability,
		Performance:  performance,
		SSL:          sslInfo,
		DNS:          dns,
		CORSDomains:  []string{},
	}, nil
}

func wrapTargetUnavailableResponse() (dto.ServiceStatsResponse, error) {
	return dto.ServiceStatsResponse{
		Availability: dto.Availability{
			IsAvailable:    false,
			HTTPStatusCode: nil,
		},
	}, ErrTargetUnavailable
}

func formatSpeed(bytesPerSecond int64) string {
	if bytesPerSecond < 1024 {
		return fmt.Sprintf("%d Bps", bytesPerSecond)
	} else if bytesPerSecond < 1024*1024 {
		return fmt.Sprintf("%.2f KBps", float64(bytesPerSecond)/1024)
	} else if bytesPerSecond < 1024*1024*1024 {
		return fmt.Sprintf("%.2f MBps", float64(bytesPerSecond)/(1024*1024))
	}
	return fmt.Sprintf("%.2f GBps", float64(bytesPerSecond)/(1024*1024*1024))
}

func getDNSDetails(domain string) dto.DNS {
	start := time.Now()
	ips, err := net.LookupHost(domain)
	if err != nil {
		return dto.DNS{}
	}
	elapsed := time.Since(start)
	return dto.DNS{
		ResponseTime: elapsed.String(),
		Servers:      ips,
	}
}
