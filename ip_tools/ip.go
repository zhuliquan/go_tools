package ip_tools

import (
	"fmt"
	"math/rand"
	"net"
)

type Ip struct {
	ipType IpType
	value  net.IP
}

func (i *Ip) String() string {
	return i.value.String()
}

func (i *Ip) GetIpType() IpType {
	return i.ipType
}

func (i *Ip) GetValue() []byte {
	return []byte(i.value)
}

func NewIp(ip string) (*Ip, error) {
	x, err := parseIp(ip)
	if err != nil {
		return nil, err
	}
	switch getIpType(ip) {
	case IPV4:
		return &Ip{value: x.To4(), ipType: IPV4}, nil
	case IPV6:
		return &Ip{value: x.To16(), ipType: IPV6}, nil
	default:
		return nil, fmt.Errorf("unknown ip type")
	}
}

type IpCidr struct {
	ipType        IpType
	broadcastAddr net.IP
	ipNet         *net.IPNet
	maskLength    int
}

func (i *IpCidr) String() string {
	return i.ipNet.String()
}

func (i *IpCidr) GetValue() IpType {
	return i.ipType
}

// get ip cidr network address
func (i *IpCidr) GetNetworkAddress() []byte {
	if i.ipType == IPV4 && (i.maskLength == 32 || i.maskLength == 31) {
		return nil
	}
	return []byte(i.ipNet.IP)
}

// get ip cidr broadcase address
func (i *IpCidr) GetBroadCastAddress() []byte {
	return []byte(i.broadcastAddr)
}

// get ip cidr mask address
func (i *IpCidr) GetMaskAddress() []byte {
	return []byte(i.ipNet.Mask)
}

// get ip cidr mask length (the number of sequence ones)
func (i *IpCidr) GetMaskLength() int {
	return i.maskLength
}

// get first valid ip
// if masklength == 31/32   && iptype == IPV4 first valid address is network address
// if masklength == 127/128 && iptype == IPV6 first valid address is network address
func (i *IpCidr) GetFirstValidIp() []byte {
	if (i.ipType == IPV4 && (i.maskLength == 31 || i.maskLength == 32)) ||
		(i.ipType == IPV6 && (i.maskLength == 127 || i.maskLength == 128)) {
		return i.ipNet.IP
	}
	return add(i.ipNet.IP, 1)
}

// get last valid ip
// if masklength == 31/32   && iptype == IPV4 first valid address is broadcast address
// if masklength == 127/128 && iptype == IPV6 first valid address is broadcast address
func (i *IpCidr) GetLastValidIp() []byte {
	if (i.ipType == IPV4 && (i.maskLength == 31 || i.maskLength == 32)) ||
		(i.ipType == IPV6 && (i.maskLength == 127 || i.maskLength == 128)) {
		return i.broadcastAddr
	}
	return dec(i.broadcastAddr, 1)
}

// get random ip between network address and broadcase address, ip ~ (netwok address, broadcase address).
func (i *IpCidr) GetRandomValidIp() []byte {
	if i.ipType == IPV4 && i.maskLength == 32 {
		return i.ipNet.IP
	} else if i.ipType == IPV6 && i.maskLength == 128 {
		return i.ipNet.IP
	} else if (i.ipType == IPV4 && i.maskLength == 31) ||
		(i.ipType == IPV6 && i.maskLength == 127) {
		if rand.Float64() < 0.5 {
			return i.ipNet.IP
		} else {
			return i.broadcastAddr
		}
	} else {
		firstValidIp := i.GetFirstValidIp()
		lastValidIp := i.GetLastValidIp()
		randomIp, _ := getBetweenRandomIp(firstValidIp, lastValidIp)
		return randomIp
	}
}

func NewIpCidr(ipCidr string) (*IpCidr, error) {
	ipNet, err := parseIpCidr(ipCidr)
	if err != nil {
		return nil, err
	}
	var ipType IpType
	broadCast := getBroadcaseAddress(ipNet)
	switch getIpType(ipCidr) {
	case IPV4:
		ipType = IPV4
		ipNet.IP = ipNet.IP.To4()
		broadCast = broadCast.To4()
	case IPV6:
		ipType = IPV6
		ipNet.IP = ipNet.IP.To16()
		broadCast = broadCast.To16()
	default:
		return nil, fmt.Errorf("unknown ip type")
	}
	return &IpCidr{
		ipType:        ipType,
		ipNet:         ipNet,
		broadcastAddr: broadCast,
		maskLength:    getMaskOnes(ipNet),
	}, nil
}
