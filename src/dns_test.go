package src

import (
	"fmt"
	"testing"
)

func TestGenerateDNSRecord(t *testing.T) {
	r := GenerateDNSRecord("2001:db8::1", 3600, "foo")
	exp := "foo    3600    IN    AAAA    2001:db8::1"

	if r != exp {
		t.Fatalf("%s != %s", r, exp)
	}
}

func TestGenerateDNSRecordFQDN(t *testing.T) {
	r := GenerateDNSRecord("2001:db8::1", 3600, "foo.bar")
	exp := "foo.bar.    3600    IN    AAAA    2001:db8::1"

	if r != exp {
		t.Fatalf("%s != %s", r, exp)
	}
}

func TestGenerateReverseDNSRecord(t *testing.T) {
	ip := "b.4.7.e.0.2.4.c.a.d.f.0.c.d.4.1.0.0.0.0.0.0.0.0.8.b.d.0.1.0.0.2.ip6.arpa"
	r := GenerateReverseDNSRecord(ip, 3600, "foo.bar")
	exp := fmt.Sprintf("%s.    3600    IN    PTR    foo.bar", ip)

	if r != exp {
		t.Fatalf("%s != %s", r, exp)
	}
}
