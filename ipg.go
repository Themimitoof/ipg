package main

import (
	"fmt"
	"net"
	"os"

	"github.com/alecthomas/kong"
	"github.com/c-robinson/iplib"
	"github.com/themimitoof/ipg/src"
	"github.com/themimitoof/ipg/src/output"
)

var cliOutputSettings output.OutputInformationSettings
var cli struct {
	Subnet        string `arg:"" help:"IPv6 Subnet"`
	Random        bool   `short:"r" required:"" help:"Generate a random IPv6 address on the given subnet."`
	Name          string `short:"n" required:"" default:"hostname" help:"Specify the hostname of a machine, an IPv6 address will be generated based on it."`
	Silent        bool   `short:"s" help:"Only display values without their labels."`
	Format        string `name:"format" short:"f" enum:"console,json" default:"console" help:"Specify the type of output. Possible values: console, json"`
	Address       bool   `short:"a" help:"Display the generated IP address."`
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

	cliOutputSettings = output.OutputInformationSettings{
		Address:       cli.Address,
		Reverse:       cli.Reverse,
		DNSRecord:     cli.DNSRecord,
		ReverseRecord: cli.ReverseRecord,
		Format:        cli.Format,
		Silent:        cli.Silent,
	}

	if !cli.Address && !cli.Reverse && !cli.DNSRecord && !cli.ReverseRecord {
		cliOutputSettings.AllData = true
	}

	var ipNetwork iplib.Net6 = iplib.Net6FromStr(cli.Subnet)

	// Check if the Subnet is valid or not.
	if len(ipNetwork.IP()) == 0 {
		os.Stderr.WriteString("The given subnet is not valid.\n")
		os.Exit(1)
	}

	// Check if the subnet is not too small to be used
	var cidr, _ = ipNetwork.Mask().Size()

	if cidr > 126 {
		os.Stderr.WriteString(
			fmt.Sprintf("The given subnet (/%d) is too small to be used with ipg.\n", cidr),
		)
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
	var dnsRecord string = src.GenerateDNSRecord(ipAddr, cli.DNSTTL, cli.Name)
	var reverseDnsRecord string = src.GenerateReverseDNSRecord(reverseIpAddr, cli.DNSTTL, cli.Name)
	var cmdOutput []byte

	// Render the output
	if cli.Format == "json" {
		cmdOutput = output.IpgJsonOutput{
			Config: &cliOutputSettings,
			Data: output.IpgOutputData{
				Hostname:      cli.Name,
				IPAddress:     generatedIp.String(),
				IPReverse:     reverseIpAddr,
				DNSRecord:     dnsRecord,
				ReverseRecord: reverseDnsRecord,
			},
		}.Render()
	} else {
		cmdOutput = output.IpgConsoleOutput{
			Config: &cliOutputSettings,
			Data: output.IpgOutputData{
				Hostname:      cli.Name,
				IPAddress:     generatedIp.String(),
				IPReverse:     iplib.IPToARPA(generatedIp),
				DNSRecord:     dnsRecord,
				ReverseRecord: reverseDnsRecord,
			},
		}.Render()
	}

	os.Stdout.Write(cmdOutput)
}
