package main

import (
	"fmt"
	"net"
	"os"

	"github.com/alecthomas/kong"
	"github.com/c-robinson/iplib"
	"github.com/themimitoof/ipg/src"
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
		generatedIp = src.GenerateRandomIP(ipNetwork)
	} else {
		generatedIp = src.GenerateIPFromHostname(ipNetwork, cli.Name)
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
