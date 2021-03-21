package ip_tools

import (
	"fmt"
	"math/rand"
	"net"
)

// return ip is valid, or not
func CheckValidIp(ip string) error {
	_, err := parseIp(ip)
	return err
}

// return ip cidr is valid, or not
func CheckValidIpCidr(ipCidr string) error {
	_, err := parseIpCidr(ipCidr)
	return err
}

// return ip type [ipv4/ipv6], before use this function, you must ensure ip is valid.
func GetValidIpType(ip string) IpType {
	return getIpType(ip)
}

// return ip cidr type [ipv4/ipv6], before use this function, you must ensure ip cidr is valid.
func GetValidIpCidrType(ipCidr string) IpType {
	return getIpType(ipCidr)
}

// ip to byte array
func Ip2Bytes(ip string) []byte {
	x, err := parseIp(ip)
	if err != nil {
		return nil
	}
	switch getIpType(ip) {
	case IPV4:
		return x.To4()
	default:
		return x.To16()
	}
}

// get a random ipv4
func GetRandomIpV4() *Ip {
	x1 := byte(rand.Intn(256))
	x2 := byte(rand.Intn(256))
	x3 := byte(rand.Intn(256))
	x4 := byte(rand.Intn(256))
	return &Ip{
		ipType: IPV4,
		value:  net.IP{x1, x2, x3, x4}.To4(),
	}
}

// get a random ipv6
func GetRandomIpV6() *Ip {
	res := make([]byte, 16)
	for i := 0; i < 16; i++ {
		res[i] = byte(rand.Intn(256))
	}
	return &Ip{
		ipType: IPV6,
		value:  net.IP(res).To16(),
	}
}

// get a random ip cidr v4
func GetRandomIpCidrV4() *IpCidr {
	maskLength := rand.Intn(30) + 1
	ipV4 := GetRandomIpV4()
	ipCidr, _ := NewIpCidr(fmt.Sprintf("%s/%d", ipV4, maskLength))
	// fmt.Println(ipCidr)
	return ipCidr
}

// get a random ip cidr v6
func GetRandomIpCidrV6() *IpCidr {
	maskLength := rand.Intn(126) + 1
	ipV6 := GetRandomIpV6()
	ipCidr, _ := NewIpCidr(fmt.Sprintf("%s/%d", ipV6, maskLength))
	// fmt.Println(ipCidr)
	return ipCidr
}

// get a first valid ip and last valid ip by ip cidr
func GetRangeIpByIpCidr(ipCidr string) ([]byte, []byte, error) {
	x, err := NewIpCidr(ipCidr)
	if err != nil {
		return nil, nil, err
	}
	return x.GetFirstValidIp(), x.GetLastValidIp(), nil

}

// get a ip cidr network address
func GetNetworkAddr(ipCidr string) ([]byte, error) {
	x, err := NewIpCidr(ipCidr)
	if err != nil {
		return nil, err
	}
	return x.GetNetworkAddress(), nil
}

// get ip cidr broadcase address
func GetBroadCastAddr(ipCidr string) ([]byte, error) {
	x, err := NewIpCidr(ipCidr)
	if err != nil {
		return nil, err
	}
	return x.GetBroadCastAddress(), nil
}

// get ip cidr random contained ip
func GetBetweenRandomIp(ipCidr string) ([]byte, error) {
	x, err := NewIpCidr(ipCidr)
	if err != nil {
		return nil, err
	}
	return x.GetRandomValidIp(), nil
}
