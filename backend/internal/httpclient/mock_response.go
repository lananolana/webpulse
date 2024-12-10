package httpclient

import (
	"github.com/lananolana/webpulse/backend/internal/dto"
	"net"
)

var mockServiceStatResponse = dto.ServiceStatsResponse{
	Status: dto.Success,
	Availability: &dto.Availability{
		IsAvailable:    Ptr(true),
		HTTPStatusCode: Ptr(200),
	},
	Performance: &dto.Performance{
		ResponseTimeMs:    Ptr(int64(123)),
		TransferSpeedKbps: Ptr(int64(56)),
		ResponseSizeKb:    Ptr(int64(12)),
		Optimization:      Ptr("gzip"),
	},
	Security: &dto.Security{
		SSL: &dto.SSL{
			Valid:     Ptr(true),
			ExpiresAt: Ptr(int64(1735824000)),
			Issuer:    Ptr("Let's Encrypt"),
		},
		CORS: &dto.CORS{
			Enabled:     Ptr(true),
			AllowOrigin: Ptr("*"),
		},
	},
	ServerInfo: &dto.ServerInfo{
		IPAddress:         Ptr(net.ParseIP("93.184.216.34")),
		WebServer:         Ptr("nginx"),
		DnsResponseTimeMs: Ptr(int64(45)),
		DnsRecords: &dto.DnsRecords{
			A:     []net.IP{net.ParseIP("93.184.216.34")},
			CNAME: Ptr("example.com"),
			MX:    []string{"mail.example.com"},
		},
	},
}
