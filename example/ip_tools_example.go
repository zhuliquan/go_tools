package main

import (
	"fmt"

	"github.com/zhuliquan/go_tools/ip_tools"
)

func main() {
	if err := ip_tools.CheckValidIp("123.0.0.1"); err == nil {
		fmt.Println("this is valid ip")
	}
	if err := ip_tools.CheckValidIp("123.0.1"); err != nil {
		fmt.Println("this is not valid ip")
	}
	if err := ip_tools.CheckValidIpCidr("126.0.0.8/4"); err != nil {
		fmt.Println("this is valid ip cidr")
	}
	if err := ip_tools.CheckValidIpCidr("126.0.0.9/800"); err != nil {
		fmt.Println("this is not valid ip cidr")
	}
	ip := ip_tools.GetRandomIpV4()
	fmt.Println("get a random ip v4: ", ip)
	ip = ip_tools.GetRandomIpV6()
	fmt.Println("get a random ip v6: ", ip)
	ipCidr := ip_tools.GetRandomIpCidrV4()
	fmt.Printf("Get a random ip cidr v4: %s \n", ipCidr)
	fmt.Printf("accroding to ip cidr %s, network address: %+v\n", ipCidr, ipCidr.GetNetworkAddress())
	fmt.Printf("accroding to ip cidr %s, broadcast address: %+v\n", ipCidr, ipCidr.GetBroadCastAddress())
	fmt.Printf("accroding to ip cidr %s, first valid address: %+v\n", ipCidr, ipCidr.GetFirstValidIp())
	fmt.Printf("accroding to ip cidr %s, last valid address: %+v\n", ipCidr, ipCidr.GetLastValidIp())
	fmt.Printf("accroding to ip cidr %s, we can get random contained ip: %+v\n", ipCidr, ipCidr.GetRandomValidIp())
	ipCidr = ip_tools.GetRandomIpCidrV6()
	fmt.Println("Get a random ip cidr v6: ", ipCidr)
	fmt.Printf("accroding to ip cidr %s, network address: %+v\n", ipCidr, ipCidr.GetNetworkAddress())
	fmt.Printf("accroding to ip cidr %s, broadcast address: %+v\n", ipCidr, ipCidr.GetBroadCastAddress())
	fmt.Printf("accroding to ip cidr %s, first valid address: %+v\n", ipCidr, ipCidr.GetFirstValidIp())
	fmt.Printf("accroding to ip cidr %s, last valid address: %+v\n", ipCidr, ipCidr.GetLastValidIp())
	fmt.Printf("accroding to ip cidr %s, we can get random contained ip: %+v\n", ipCidr, ipCidr.GetRandomValidIp())
}
