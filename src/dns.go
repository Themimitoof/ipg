package src

import (
	"fmt"
	"strings"
)

// Generate a DNS record ready to be paste on a Bind compatible zone
func GenerateDNSRecord(ip string, ttl int, hostname string) string {
	if strings.Contains(hostname, ".") && hostname[len(hostname)-1] != '.' {
		hostname += "."
	}

	return fmt.Sprintf("%s    %d    IN    AAAA    %s", hostname, ttl, ip)
}

// Generate a ARPA record ready to be paste on a Bind compatible zone
func GenerateReverseDNSRecord(ip string, ttl int, hostname string) string {
	return fmt.Sprintf("%s.    %d    IN    PTR    %s", ip, ttl, hostname)
}
