package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
type Stringer interface {
	String() string
}

func (ipAddress IPAddr) String() string {
	const FULLSTOP_SEPERATOR = ".";
	ipAddressString := string(ipAddress[0]) + FULLSTOP_SEPERATOR + string(ipAddress[1]) + FULLSTOP_SEPERATOR + string(ipAddress[2]) + FULLSTOP_SEPERATOR + string(ipAddress[3]);
	return ipAddressString;
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
