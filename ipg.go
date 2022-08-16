package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"

	"github.com/alecthomas/kong"
	"github.com/c-robinson/iplib"
	"github.com/fatih/color"
)

var cli struct {
	Subnet        string `arg:"" help:"IPv6 Subnet"`
	Random        bool   `short:"r" xor:"Name,Random" required:"" help:"Generate a random IPv6 address on the given subnet."`
	Name          string `short:"n" xor:"Name,Random" required:"" help:"Specify the hostname of a machine, an IPv6 address will be generated based on it."`
	Silent        bool   `short:"s" help:"Only display values without their labels."`
	Reverse       bool   `short:"R" help:"Display the ARPA version of the IP address."`
	DNSRecord     bool   `name:"dns" short:"d" help:"Returns a DNS record ready to paste to a DNS zone."`
	ReverseRecord bool   `name:"rrecord" short:"x" help:"Returns a rDNS record ready to paste to a DNS zone."`
	DNSTTL        int    `name:"ttl" short:"t" default:"86400" help:"TTL value for DNS returned DNS records."`
}

// Function stollen from c-robinson/iplib, method non-exposed by the lib ¯\_(ツ)_/¯
func Net6wildcard(n iplib.Net6) net.IPMask {
	wc := make([]byte, len(n.Mask()))
	for i, b := range n.Mask() {
		wc[i] = 0xff - b
	}
	return wc
}

func RandomNumber(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
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

// Generate a DNS record ready to be paste on a Bind compatible zone
func GenerateDNSRecord(ip string, ttl int, hostname string) string {
	return fmt.Sprintf("%s\t%d\tIN\tAAAA\t%s", hostname, ttl, ip)
}

// Generate a ARPA record ready to be paste on a Bind compatible zone
func GenerateReverseDNSRecord(ip string, ttl int, hostname string) string {
	return fmt.Sprintf("%s.\t%d\tIN\tPTR\t%s", ip, ttl, hostname)
}

func main() {
	kong.Parse(
		&cli,
		kong.Description("A simple IPv6 generator for lazy netadmins."),
	)

	var ipNetwork iplib.Net6 = iplib.Net6FromStr(cli.Subnet)

	// Check if the Subnet is valid or not.
	if len(ipNetwork.IP()) == 0 {
		fmt.Println("The given subnet is not valid.")
		os.Exit(1)
	}

	// Check if the subnet is not too small to be used
	var cidr, _ = ipNetwork.Mask().Size()

	if cidr > 126 {
		fmt.Printf("The given subnet is too small (/%d) to be used with ipg.\n", cidr)
		os.Exit(1)
	}

	// Generate the new IP address using the given generation method
	var generatedIp net.IP

	if cli.Random {
		generatedIp = GenerateRandomIP(ipNetwork)
	} else {
		generatedIp = GenerateIPFromHostname(ipNetwork, cli.Name)
	}

	var reverseIpAddr string = iplib.IPToARPA(generatedIp)
	var ipAddr string = generatedIp.String()

	// Always display the generated IP address
	if !cli.Silent {
		fmt.Printf("IP address: ")
	}
	color.Yellow(ipAddr)

	// Display the ARPA version of the IP address if the flag is set
	if cli.Reverse {
		if !cli.Silent {
			fmt.Printf("Reverse IP address: ")
		}
		color.Yellow(reverseIpAddr)
	}

	// Display the BIND DNS record if the flag is set
	if cli.DNSRecord {
		if !cli.Silent {
			fmt.Println("DNS record:")
		}
		color.Yellow(GenerateDNSRecord(ipAddr, cli.DNSTTL, cli.Name))
	}

	// Display the ARPA DNS record if the flag is set
	if cli.ReverseRecord {
		if !cli.Silent {
			fmt.Println("ARPA DNS record:")
		}
		color.Yellow(GenerateReverseDNSRecord(reverseIpAddr, cli.DNSTTL, cli.Name))
	}
}
