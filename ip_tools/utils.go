package ip_tools

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
)

// parse ip by net package
func parseIp(ip string) (net.IP, error) {
	if len(ip) == 0 {
		return nil, fmt.Errorf("ip is empty")
	} else if x := net.ParseIP(ip); x == nil {
		return nil, fmt.Errorf("ip: '%s' format is incorrect", ip)
	} else {
		return x, nil
	}
}

// parse ip cidr by net package
func parseIpCidr(ipCidr string) (*net.IPNet, error) {
	if len(ipCidr) == 0 {
		return nil, fmt.Errorf("ip cidr is empty")
	} else if _, ipNet, err := net.ParseCIDR(ipCidr); err != nil {
		return nil, fmt.Errorf("ip cidr: '%s' format is incorrect", ipCidr)
	} else {
		return ipNet, nil
	}
}

// get ip type by whether it contain ':' / '.'
// if it contain ':', will return ipv6
// if it contain '.'. will return ipv4
func getIpType(ip string) IpType {
	if strings.ContainsAny(ip, ":") {
		return IPV6
	} else if strings.ContainsAny(ip, ".") {
		return IPV4
	} else {
		return UNKNOWN
	}
}

// get broadcast according to network ip (network address and mask ip)
func getBroadcaseAddress(ipNet *net.IPNet) net.IP {
	broadcastAddr := make([]byte, len(ipNet.IP))
	for i := 0; i < len(ipNet.IP); i++ {
		broadcastAddr[i] = ipNet.IP[i] + (byte(255) - ipNet.Mask[i])
	}
	return net.IP(broadcastAddr)
}

// get ip mask ones lengths
func getMaskOnes(ipNet *net.IPNet) int {
	ones, _ := ipNet.Mask.Size()
	return ones
}

// add x (byte) to s ([]byte), consider carry number
func add(s []byte, x byte) []byte {
	c := int(x)
	rs := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ints := int(s[i])
		rs[len(s)-1-i] = byte((ints + c) % 256)
		c = (ints + c) / 256
	}
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-i-1] = rs[len(rs)-i-1], rs[i]
	}
	return rs
}

// dec x (byte) from x ([]byte), consider carry number
func dec(s []byte, x byte) []byte {
	d := int(x)
	rs := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		ints := int(s[i])
		if ints-d < 0 {
			rs[len(s)-1-i] = byte(256 + ints - d)
			d = 1
		} else {
			rs[len(s)-1-i] = byte(ints - d)
			d = 0
		}
	}
	for i := 0; i < len(rs)/2; i++ {
		rs[i], rs[len(rs)-i-1] = rs[len(rs)-i-1], rs[i]
	}
	return rs
}

// get random ip between [s, e]
func getBetweenRandomIp(s []byte, e []byte) ([]byte, error) {
	if len(s) != len(e) {
		return nil, fmt.Errorf("length of start ip is not equal end ip")
	}
	res := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == e[i] {
			res[i] = s[i]
		} else {
			res[i] = byte(rand.Intn(int(e[i])-int(s[i])+1) + int(s[i]))
		}
	}
	return res, nil
}
