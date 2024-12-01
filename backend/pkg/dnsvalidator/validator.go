package dnsvalidator

import (
	"regexp"
)

const (
	dnsName string = `^(?i)[a-z0-9а-яё](?:[a-z0-9а-яё-]*[a-z0-9а-яё])?(\.[a-z0-9а-яё](?:[a-z0-9а-яё-]*[a-z0-9а-яё])?)*\.[a-zа-яё]{2,}$`
)

var (
	dnsNameRegex = regexp.MustCompile(dnsName)
)

func DomainIsValid(domain string) bool {
	return dnsNameRegex.MatchString(domain)
}
