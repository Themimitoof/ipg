package output

import (
	"bytes"
	"fmt"

	"github.com/fatih/color"
)

type IpgConsoleOutput struct {
	Config *OutputInformationSettings
	Data   IpgOutputData
}

func (c IpgConsoleOutput) Render() []byte {
	output := bytes.Buffer{}

	// Display the generated IP address
	if c.Config.AllData || c.Config.Address {
		if c.Config.Silent {
			output.WriteString(fmt.Sprintf("%s\n", c.Data.IPAddress))
		} else {
			output.WriteString("IP address: ")
			output.WriteString(color.YellowString("%s\n", c.Data.IPAddress))
		}
	}

	// Display the ARPA version of the IP address if the flag is set
	if c.Config.AllData || c.Config.Reverse {
		if c.Config.Silent {
			output.WriteString(fmt.Sprintf("%s\n", c.Data.IPReverse))
		} else {
			output.WriteString("Reverse IP address: ")
			output.WriteString(color.YellowString("%s\n", c.Data.IPReverse))
		}
	}

	// Display the BIND DNS record if the flag is set
	if c.Config.AllData || c.Config.DNSRecord {
		if c.Config.Silent {
			output.WriteString(fmt.Sprintf("%s\n", c.Data.DNSRecord))
		} else {
			output.WriteString("DNS record: ")
			output.WriteString(color.YellowString("%s\n", c.Data.DNSRecord))
		}
	}

	// Display the ARPA DNS record if the flag is set
	if c.Config.AllData || c.Config.ReverseRecord {
		if c.Config.Silent {
			output.WriteString(fmt.Sprintf("%s\n", c.Data.ReverseRecord))
		} else {
			output.WriteString("ARPA DNS record: ")
			output.WriteString(color.YellowString("%s\n", c.Data.ReverseRecord))
		}
	}

	return output.Bytes()
}
