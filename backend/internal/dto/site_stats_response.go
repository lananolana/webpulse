package dto

import "time"

// ServiceStatsResponse represents the overall site status response.
type ServiceStatsResponse struct {
	Target       string       `json:"target,omitempty"`
	Availability Availability `json:"availability,omitempty"`
	Performance  Performance  `json:"performance,omitempty"`
	SSL          SSL          `json:"ssl,omitempty"`
	DNS          DNS          `json:"dns,omitempty"`
	CORSDomains  []string     `json:"cors_domains,omitempty"`
}

// Availability contains information about site availability.
type Availability struct {
	IsAvailable    bool `json:"is_available,omitempty"`
	HTTPStatusCode *int `json:"http_status_code,omitempty"`
}

// Performance contains metrics related to site performance.
type Performance struct {
	ResponseTime  string `json:"response_time,omitempty"`
	DownloadSpeed string `json:"download_speed,omitempty"`
}

// SSL provides information about the site's SSL certificate.
type SSL struct {
	IsValid              bool      `json:"is_valid,omitempty"`
	ExpiresAt            time.Time `json:"expires_at,omitempty"`
	Issuer               string    `json:"issuer,omitempty"`
	CertificateAuthority string    `json:"certificate_authority,omitempty"`
}

// DNS contains information about dns servers and response times.
type DNS struct {
	ResponseTime string   `json:"response_time,omitempty"`
	Servers      []string `json:"servers,omitempty"`
}
