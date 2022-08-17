# ipg: A simple IPv6 address generator for lazy netadmins

`ipg` is a small tool for netadmins wanting to generate random IPv6 addresses or based on the
hostname of a machine and want to obtain the DNS and PTR records ready to paste on their DNS zone.

`ipg` is a replacement of an old script I created in 2020 called
[v6gen](https://github.com/Themimitoof/v6gen).

## Features

 - Generate a random IPv6 address
 - Generate an IPv6 based on the given hostname
 - Return the ARPA version of the IPv6 address
 - Return a DNS and/or a PTR record
 - Get the output in a JSON format or in a more humain friendly one

Last remaining things to do:

 - [ ] Add some unit tests
 - [ ] Add goreleaser to the project
 - [ ] Complete this file
 - [ ] Release the first release


## Installation

Until no releases are available on GitHub, ensure you have [Go](https://go.dev/) installed and well configured. To install it, simply type:

```bash
go install github.com/themimitoof/ipg
```

Otherwise, you can use the Docker image by typing:

```bash
docker run --rm themimitoof/ipg --help
```

## Usage

```bash
$ ipg --help
Usage: ipg <subnet>

A simple IPv6 generator for lazy netadmins.

Arguments:
  <subnet>    IPv6 Subnet

Flags:
  -h, --help                Show context-sensitive help.
  -r, --random              Generate a random IPv6 address on the given subnet.
  -n, --name="hostname"     Specify the hostname of a machine, an IPv6 address will be generated based on it.
  -s, --silent              Only display values without their labels.
  -f, --format="console"    Specify the type of output. Possible values: console, json
  -a, --address             Display the generated IP address.
  -R, --reverse             Display the ARPA version of the IP address.
  -d, --dns                 Returns a DNS record ready to paste to a DNS zone.
  -x, --rrecord             Returns a rDNS record ready to paste to a DNS zone.
  -t, --ttl=86400           TTL value for DNS returned DNS records.
```

To generate a random IPv6 address:

```bash
$ ipg -r 2001:db8:beef::/64
IP address: 2001:db8:beef:0:9e20:7abf:4b9c:2d45
Reverse IP address: 5.4.d.2.c.9.b.4.f.b.a.7.0.2.e.9.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa
DNS record: hostname    86400    IN    AAAA    2001:db8:beef:0:9e20:7abf:4b9c:2d45
ARPA DNS record: 5.4.d.2.c.9.b.4.f.b.a.7.0.2.e.9.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa.    86400    IN    PTR    hostname
```

To generate an IPv6 address based on the hostname of a machine:

```bash
$ ipg -n hello 2001:db8:beef::/64
IP address: 2001:db8:beef::68:656c:6c6f
Reverse IP address: f.6.c.6.c.6.5.6.8.6.0.0.0.0.0.0.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa
DNS record: hello    86400    IN    AAAA    2001:db8:beef::68:656c:6c6f
ARPA DNS record: f.6.c.6.c.6.5.6.8.6.0.0.0.0.0.0.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa.    86400    IN    PTR    hello
```

You can also get a random IPv6 address and give the hostname of the machine to obtain a ready to paste DNS record:

```bash
$ ipg -rn hello.foobar.com 2001:db8:beef::/64
IP address: 2001:db8:beef:0:9639:1d78:7e2a:9646
Reverse IP address: 6.4.6.9.a.2.e.7.8.7.d.1.9.3.6.9.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa
DNS record: hello.foobar.com.    86400    IN    AAAA    2001:db8:beef:0:9639:1d78:7e2a:9646
ARPA DNS record: 6.4.6.9.a.2.e.7.8.7.d.1.9.3.6.9.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa.    86400    IN    PTR    hello.foobar.com
```

If you want to obtain one specific or multiple information like the IPv6 address and the reverse version:

```bash
$ ipg -rRa  2001:db8:beef::/64
IP address: 2001:db8:beef:0:d5b:ca60:88be:bfd1
Reverse IP address: 1.d.f.b.e.b.8.8.0.6.a.c.b.5.d.0.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa
```

If you want to get an information without its label (e.g for piping with another command), use `-s` flag:

```bash
$ ipg -ras  2001:db8:beef::/64
2001:db8:beef:0:d5b:ca60:88be:bfd1
```

If you want to get a JSON version of the output:

```bash
$ ipg -r -n hello.foobar.com -f json 2001:db8:beef::/64 | jq
{
  "hostname": "hello.foobar.com",
  "ip_addr": "2001:db8:beef:0:5462:5f99:1a46:a2bc",
  "arpa_addr": "c.b.2.a.6.4.a.1.9.9.f.5.2.6.4.5.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa",
  "dns_record": "hello.foobar.com.    86400    IN    AAAA    2001:db8:beef:0:5462:5f99:1a46:a2bc",
  "ptr_record": "c.b.2.a.6.4.a.1.9.9.f.5.2.6.4.5.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa.    86400    IN    PTR    hello.foobar.com"
}
```

**Note:** filters are not available with the `json` format.

## License

This project is released under the [MIT license](LICENSE). Feel free to use, contribute, fork and do
what you want with it. Please keep all licenses, copyright notices and mentions in case you use,
re-use, steal, fork code from this repository.
