package output

type OutputInformationSettings struct {
	AllData       bool
	Address       bool
	Reverse       bool
	DNSRecord     bool
	ReverseRecord bool

	Format string
	Silent bool
}

type IpgOutputData struct {
	Hostname      string
	IPAddress     string
	IPReverse     string
	DNSRecord     string
	ReverseRecord string
}
