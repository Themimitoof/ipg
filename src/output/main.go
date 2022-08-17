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
	Hostname      string `json:"hostname"`
	IPAddress     string `json:"ip_addr"`
	IPReverse     string `json:"arpa_addr"`
	DNSRecord     string `json:"dns_record"`
	ReverseRecord string `json:"ptr_record"`
}
