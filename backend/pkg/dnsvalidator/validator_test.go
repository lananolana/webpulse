package dnsvalidator

import (
	"testing"
)

func TestDomainIsValid(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{"Valid domain", "example.com", true},
		{"Valid domain with subdomain", "sub.example.com", true},
		{"Valid Cyrillic domain", "пример.рф", true},
		{"Invalid domain without TLD", "example", false},
		{"Invalid domain with space", "example .com", false},
		{"Invalid domain with special chars", "exa$mple.com", false},
		{"Invalid domain with trailing dash", "example-.com", false},
		{"Invalid domain with double dots", "example..com", false},
		{"Empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DomainIsValid(tt.domain)
			if result != tt.expected {
				t.Errorf("domainIsValid(%q) = %v; want %v", tt.domain, result, tt.expected)
			}
		})
	}
}

func BenchmarkDomainIsValid(b *testing.B) {
	domains := []string{
		"example.com",
		"sub.example.com",
		"пример.рф",
		"xn--d1acufc.xn--p1ai",
		"example",
		"exa$mple.com",
		"example-.com",
		"example..com",
		"",
	}

	for _, domain := range domains {
		b.Run(domain, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				DomainIsValid(domain)
			}
		})
	}
}
