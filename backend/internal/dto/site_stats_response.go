package dto

import "net"

type Status string

const (
	Success Status = "Success"
	Failed  Status = "Failed"
)

// ServiceStatsResponse represents the overall site status response.
type ServiceStatsResponse struct {
	Status       Status        `json:"status"`
	Availability *Availability `json:"availability,omitempty"`
	Performance  *Performance  `json:"performance,omitempty"`
	Security     *Security     `json:"security,omitempty"`
	ServerInfo   *ServerInfo   `json:"server_info,omitempty"`
}

// Availability contains information about site availability.
type Availability struct {
	IsAvailable    *bool `json:"is_available,omitempty"`
	HTTPStatusCode *int  `json:"http_status_code,omitempty"`
}

// Performance contains metrics related to site performance.
type Performance struct {
	ResponseTimeMs    *int64  `json:"response_time_ms,omitempty"`
	TransferSpeedKbps *int64  `json:"transfer_speed_kbps,omitempty"`
	ResponseSizeKb    *int64  `json:"response_size_kb,omitempty"`
	Optimization      *string `json:"optimization,omitempty"`
}

type Security struct {
	SSL  *SSL  `json:"ssl,omitempty"`
	CORS *CORS `json:"cors,omitempty"`
}

type ServerInfo struct {
	IPAddress         *net.IP     `json:"ip_address,omitempty"`
	WebServer         *string     `json:"web_server,omitempty"`
	DnsResponseTimeMs *int64      `json:"dns_response_time_ms,omitempty"`
	DnsRecords        *DnsRecords `json:"dns_records,omitempty"`
}

// SSL provides information about the site's SSL certificate.
type SSL struct {
	Valid     *bool   `json:"valid,omitempty"`
	ExpiresAt *int64  `json:"expires_at,omitempty"`
	Issuer    *string `json:"issuer,omitempty"`
}

type CORS struct {
	Enabled     *bool   `json:"enabled,omitempty"`
	AllowOrigin *string `json:"allow_origin,omitempty"`
}

type DnsRecords struct {
	A     []net.IP `json:"A,omitempty"`
	CNAME *string  `json:"CNAME,omitempty"`
	MX    []string `json:"MX,omitempty"`
}
