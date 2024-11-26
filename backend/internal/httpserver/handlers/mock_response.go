package handlers

const mockResponseStatus = `
{
    "status": "success",
    "availability": {
        "is_available": true,
        "http_status": 200
    },
    "performance": {
        "response_time_ms": 123,
        "transfer_speed_kbps": 56.78,
        "response_size_kb": 12.34,
        "optimization": "gzip"
    },
    "security": {
        "ssl": {
            "valid": true,
            "expires_in_timestamp": 1735824000,
            "issuer": "Let's Encrypt"
        },
        "cors": {
            "enabled": true,
            "allow_origin": "*"
        }
    },
    "server_info": {
        "ip_address": "93.184.216.34",
        "web_server": "nginx",
        "dns_response_time_ms": 45,
        "dns_records": {
            "A": [
                "93.184.216.34"
            ],
            "CNAME": [],
            "MX": [
                "mail.example.com"
            ]
        }
    }
}
`
