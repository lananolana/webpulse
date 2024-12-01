package httpclient

import (
	"github.com/lananolana/webpulse/backend/internal/dto"
	"net"
	"net/http"
	"net/url"
)

// parseURL parses the domain into a URL.
func parseURL(domain string) (*url.URL, error) {
	u, err := url.Parse(domain)
	if err != nil {
		return nil, err
	}
	if u.Host == "" {
		u.Host = domain
	}
	return u, nil
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
		IsAvailable:    Ptr(true),
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
		AllowOrigin: Ptr(allowOrigin),
		Enabled:     Ptr(enabled),
	}
}

func getDNSDetails(domain string) *dto.DnsRecords {
	var dnsRecords dto.DnsRecords

	ips, err := net.LookupIP(domain)
	if err == nil {
		dnsRecords.A = ips
	}

	mxs, err := net.LookupMX(domain)
	if err == nil {
		mxDomains := make([]string, len(mxs))
		for i, mx := range mxs {
			mxDomains[i] = mx.Host
		}
		dnsRecords.MX = mxDomains
	}

	cname, err := net.LookupCNAME(domain)
	if err == nil {
		dnsRecords.CNAME = Ptr(cname)
	}

	return &dnsRecords
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
