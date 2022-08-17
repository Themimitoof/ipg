package src

import (
	"net"
	"strings"

	"github.com/c-robinson/iplib"
)

// Function stollen from c-robinson/iplib, method non-exposed by the lib ¯\_(ツ)_/¯
func Net6wildcard(n iplib.Net6) net.IPMask {
	wc := make([]byte, len(n.Mask()))
	for i, b := range n.Mask() {
		wc[i] = 0xff - b
	}
	return wc
}

// Generate a random IP address based on the given IP subnet.
func GenerateRandomIP(net iplib.Net6) net.IP {
	finalIP := net.FirstAddress()

	for idx, b := range net.Mask() {
		if b == 0xff {
			continue
		} else if b != 0x0 {
			randByte := byte(RandomNumber(0, 255))
			wildCardByte := Net6wildcard(net)[idx]

			finalIP[idx] = net.IP()[idx] + wildCardByte&randByte
		} else {
			finalIP[idx] = byte(RandomNumber(0, 255))
		}
	}

	return finalIP
}

// Generate an IP address based on the hostname.
// The algorithm first tries with the complete hostname, if it does not work,
// it tries by removing vowels and iterate by removing character by character.
func GenerateIPFromHostname(net iplib.Net6, hostname string) net.IP {
	finalIP := []byte(net.FirstAddress())
	cidr, _ := net.Mask().Size()
	availableBits := 128 - cidr
	hostname = strings.Split(hostname, ".")[0]

	if len(hostname)*8 > availableBits {
		// Remove vowels
		shrinkedHostname := []byte{}
		for _, c := range hostname {
			asVowel := false

			for _, v := range "aiueo" {
				if c == v {
					asVowel = true
				}
			}

			if !asVowel {
				shrinkedHostname = append(shrinkedHostname, byte(c))
			}
		}

		// In case it's not enough, shrink last character until it fits
		for len(shrinkedHostname)*8 > availableBits {
			shrinkedHostname = shrinkedHostname[0 : len(shrinkedHostname)-1]
		}

		hostname = string(shrinkedHostname)
	}

	for i, c := range hostname {
		pos := len(finalIP) - len(hostname) + i
		finalIP[pos] = byte(c)
	}

	return finalIP
}
