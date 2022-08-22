package src

import (
	"bytes"
	"net"
	"testing"

	"github.com/c-robinson/iplib"
)

func TestNet6wildcard(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0::/64")
	wildcard := Net6wildcard(n)
	exp := net.IPMask{0, 0, 0, 0, 0, 0, 0, 0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}

	if bytes.Compare(wildcard, exp) != 0 {
		t.Fatalf("The wildcard doesn't match with the expected one.")
	}
}

func TestGenerateRandomIP(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0::/64")
	ip := GenerateRandomIP(n)
	exp := []byte{0x20, 0x01, 0x0d, 0xb8, 0xba, 0xba, 0xb0, 0xb0}

	if bytes.Compare(ip[0:8], exp) != 0 {
		t.Fatalf("%s is outside of %s", ip.String(), n.String())
	}
}

func TestGenerateRandomIPKeepByteInSmallerSubnets(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0:ee::/79")
	ip := GenerateRandomIP(n)

	if ip[9] != 0xee && ip[9] != 0xef {
		t.Fatalf(
			"Last bytes of the netpart in the generated IP has shifted. Net: %s, Generated IP: %s",
			n.String(),
			ip.String(),
		)
	}
}

func TestGenerateIPFromHostname(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0::/64")
	ip := GenerateIPFromHostname(n, "foo.bar")

	if bytes.Compare(ip[13:16], []byte{0x66, 0x6f, 0x6f}) != 0 {
		t.Fatalf("'foo' as not been found in the generated IP address %s.", ip.String())
	}
}

func TestGenerateIPFromHostnameRemoveVowels(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0::/64")
	ip := GenerateIPFromHostname(n, "hellofolks.foo.bar") // hllflks

	if bytes.Compare(ip[9:16], []byte{0x68, 0x6c, 0x6c, 0x66, 0x6c, 0x6b, 0x73}) != 0 {
		t.Fatalf("'hllflks' as not been found in the generated IP address %s.", ip.String())
	}
}

func TestGenerateIPFromHostnameShrink(t *testing.T) {
	n := iplib.Net6FromStr("2001:db8:baba:b0b0::/64")
	ip := GenerateIPFromHostname(n, "loadbalancer-pouet.foo.bar") // ldblncr-

	if bytes.Compare(ip[8:16], []byte{0x6c, 0x64, 0x62, 0x6c, 0x6e, 0x63, 0x72, 0x2d}) != 0 {
		t.Fatalf("'ldblncr-' as not been found in the generated IP address %s.", ip.String())
	}
}
