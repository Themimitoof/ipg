# ipg: A simple IPv6 address generator for lazy netadmins

...

`ipg` is a replacement of an old script I created called [v6gen](https://github.com/Themimitoof/v6gen).

```bash
$ ipg --help
Usage: ipg --random --name=STRING <subnet>

A simple IPv6 generator for lazy netadmins.

Arguments:
  <subnet>    IPv6 Subnet

Flags:
  -h, --help           Show context-sensitive help.
  -r, --random         Generate a random IPv6 address on the given subnet.
  -n, --name=STRING    Specify the hostname of a machine, an IPv6 address will be generated based on it.
  -s, --silent         Only display values without their labels.
  -R, --reverse        Display the ARPA version of the IP address.
  -d, --dns            Returns a DNS record ready to paste to a DNS zone.
  -x, --rrecord        Returns a rDNS record ready to paste to a DNS zone.
  -t, --ttl=86400      TTL value for DNS returned DNS records.
```

## Features

 - Generate a random IPv6 address
 - Generate an IPv6 based on the given hostname
 - Return the ARPA version of the IPv6 address
 - Return a DNS and/or a PTR record

## Things to do

 - Add some unit tests
 - Restructure a little bit the code
 - Add the possibility to choose the `json`, `console` or something else as output
 - Create a minimal docker image (with [Busybox](https://hub.docker.com/_/busybox))
 - Add goreleaser to the project
 - Complete this file
 - Release the first release

## Examples


IP generated from an hostname, with the DNS and the PTR record:
```bash
$ ipg -n pouet.foobar.corp 2001:db8:beef::/64 -Rdx
IP address: 2001:db8:beef::70:6f75:6574
Reverse IP address: 4.7.5.6.5.7.f.6.0.7.0.0.0.0.0.0.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa
DNS record:
pouet.foobar.corp       86400   IN      AAAA    2001:db8:beef::70:6f75:6574
ARPA DNS record:
4.7.5.6.5.7.f.6.0.7.0.0.0.0.0.0.0.0.0.0.f.e.e.b.8.b.d.0.1.0.0.2.ip6.arpa.       86400   IN      PTR     pouet.foobar.corp
```

IP randomly generated:
```bash
$ ipg -r 2001:db8:beef::/64
IP address: 2001:db8:beef:0:ad15:f11f:2ec7:7268
```

Only returning a randomly generated IPv6 (useful for piping or storing the value in a shell variable):
```bash
$ ipg -sr 2001:db8:beef::/64
2001:db8:beef:0:10c2:76f4:f41:eee0
```

## License

This project is released under the [MIT license](LICENSE). Feel free to use, contribute, fork and do
what you want with it. Please keep all licenses, copyright notices and mentions in case you use,
re-use, steal, fork code from this repository.
