package handlers

import (
	"context"
	"github.com/lananolana/webpulse/backend/internal/dto"
	"github.com/lananolana/webpulse/backend/internal/httpclient"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestSiteHandler_GetSiteStatus(t *testing.T) {

	var (
		exampleSuccessResponse = dto.ServiceStatsResponse{
			Status: dto.Success,
			Msg:    nil,
			Availability: &dto.Availability{
				IsAvailable:    httpclient.Ptr(true),
				HTTPStatusCode: httpclient.Ptr(200),
			},
			Performance: &dto.Performance{
				ResponseTimeMs:    httpclient.Ptr(int64(123)),
				TransferSpeedKbps: httpclient.Ptr(int64(56)),
				ResponseSizeKb:    httpclient.Ptr(int64(12)),
				Optimization:      httpclient.Ptr("gzip"),
			},
			Security: &dto.Security{
				SSL: &dto.SSL{
					Valid:     httpclient.Ptr(true),
					ExpiresAt: httpclient.Ptr(int64(1735824000)),
					Issuer:    httpclient.Ptr("Let's Encrypt"),
				},
				CORS: &dto.CORS{
					Enabled:     httpclient.Ptr(true),
					AllowOrigin: httpclient.Ptr("*"),
				},
			},
			ServerInfo: &dto.ServerInfo{
				IPAddress:         httpclient.Ptr(net.ParseIP("93.184.216.34")),
				WebServer:         httpclient.Ptr("nginx"),
				DnsResponseTimeMs: httpclient.Ptr(int64(45)),
				DnsRecords: &dto.DnsRecords{
					A:     []net.IP{net.ParseIP("93.184.216.34")},
					CNAME: httpclient.Ptr("example.com"),
					MX:    []string{"mail.example.com"},
				},
			},
		}

		exampleResponseInvalidDomain = dto.ServiceStatsResponse{
			Status: dto.Failed,
			Msg:    httpclient.Ptr("invalid domain"),
		}
	)

	tests := []struct {
		name                   string
		requestDomain          string
		getServiceStatCallMock dto.ServiceStatsResponse
		wantResponse           dto.ServiceStatsResponse
	}{
		{
			name:                   "При успешном запросе возвращается сервисная статистика",
			requestDomain:          "google.com",
			getServiceStatCallMock: exampleSuccessResponse,
			wantResponse:           exampleSuccessResponse,
		},
		{
			name:                   "При невалидном домене ошибка возвращается в поле msg",
			requestDomain:          "invalid.co!m",
			getServiceStatCallMock: exampleResponseInvalidDomain,
			wantResponse:           exampleResponseInvalidDomain,
		},
	}

	mockSiteClient := NewMockSiteClient(t)
	ctx := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mockSiteClient.On("GetServiceStats", ctx, tt.requestDomain).Return(tt.getServiceStatCallMock)

			resp := mockSiteClient.GetServiceStats(ctx, tt.requestDomain)

			assert.Equal(t, tt.getServiceStatCallMock, resp)
		})
	}
}
