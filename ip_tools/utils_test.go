package ip_tools

import (
	// "bytes"
	"bytes"
	"net"
	"reflect"
	"testing"
)

func TestParseIp01(t *testing.T) {
	expectIp := net.IP{127, 0, 0, 1}
	nowIp, err := parseIp("127.0.0.1")
	if err != nil {
		t.Fail()
	} else if !reflect.DeepEqual(expectIp.To4(), nowIp.To4()) {
		t.Fail()
	}
}

func TestParseIp02(t *testing.T) {
	_, err := parseIp("1270.0.0.1")
	if err == nil {
		t.Fail()
	}
}

func TestParseIp03(t *testing.T) {
	_, err := parseIp("x.8.9.0")
	if err == nil {
		t.Fail()
	}
}

func TestParseIp04(t *testing.T) {
	_, err := parseIp("0.0.0.0")
	if err != nil {
		t.Fail()
	}
}

func TestParseIp05(t *testing.T) {
	_, err := parseIp("255.255.255.255")
	if err != nil {
		t.Fail()
	}
}

func TestParseIp06(t *testing.T) {
	_, err := parseIp("255.255.255")
	if err == nil {
		t.Fail()
	}
}

func TestParseIpCidr01(t *testing.T) {
	ipNet, err := parseIpCidr("127.0.0.2/24")
	if err != nil {
		t.Fail()
	} else if !reflect.DeepEqual(ipNet.IP.To4(), net.IP{127, 0, 0, 0}.To4()) {
		t.Fail()
	}
}

func TestParseIpCidr02(t *testing.T) {
	ipNet, err := parseIpCidr("127.0.0.200/24")
	if err != nil {
		t.Fail()
	} else if !reflect.DeepEqual(ipNet.IP.To4(), net.IP{127, 0, 0, 0}.To4()) {
		t.Fail()
	}
}

func TestParseIpCidr03(t *testing.T) {
	_, err := parseIpCidr("127.0.0.2")
	if err == nil {
		t.Fail()
	}
}

func TestParseIpCidr04(t *testing.T) {
	_, err := parseIpCidr("127.0.0.1/31")
	if err != nil {
		t.Fail()
	}
}

func TestParseIpCidr05(t *testing.T) {
	_, err := parseIpCidr("127.0.0.1/32")
	if err != nil {
		t.Fail()
	}
}

func TestGetBroadcaseAddress01(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.0.90/26")
	if err != nil {
		t.Fail()
	}
	ba := getBroadcaseAddress(ipNet)
	if !reflect.DeepEqual(ba.To4(), net.IP{10, 0, 0, 127}.To4()) {
		t.Fail()
	}
}

func TestGetBroadcaseAddress02(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.89.90/3")
	if err != nil {
		t.Fail()
	}
	ba := getBroadcaseAddress(ipNet)
	if !reflect.DeepEqual(ba.To4(), net.IP{31, 255, 255, 255}.To4()) {
		t.Fail()
	}
}

func TestGetBroadcaseAddress03(t *testing.T) {
	ipNet, err := parseIpCidr("127.0.0.200/24")
	if err != nil {
		t.Fail()
	}
	ba := getBroadcaseAddress(ipNet)
	if !reflect.DeepEqual(ba.To4(), net.IP{127, 0, 0, 255}.To4()) {
		t.Fail()
	}
}

func TestGetBroadcaseAddress04(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.89.90/13")
	if err != nil {
		t.Fail()
	}
	ba := getBroadcaseAddress(ipNet)
	if !reflect.DeepEqual(ba.To4(), net.IP{10, 7, 255, 255}.To4()) {
		t.Fail()
	}
}

func TestGetBroadcaseAddress05(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.230.90/23")
	if err != nil {
		t.Fail()
	}
	ba := getBroadcaseAddress(ipNet)
	if !reflect.DeepEqual(ba.To4(), net.IP{10, 0, 231, 255}.To4()) {
		t.Fail()
	}
}

func TestGetMaskOnes01(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.230.90/23")
	if err != nil {
		t.Fail()
	}
	ones := getMaskOnes(ipNet)
	if !reflect.DeepEqual(ones, 23) {
		t.Fail()
	}
}

func TestGetMaskOnes02(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.230.90/3")
	if err != nil {
		t.Fail()
	}
	ones := getMaskOnes(ipNet)
	if !reflect.DeepEqual(ones, 3) {
		t.Fail()
	}
}

func TestGetMaskOnes03(t *testing.T) {
	ipNet, err := parseIpCidr("10.0.230.90/24")
	if err != nil {
		t.Fail()
	}
	ones := getMaskOnes(ipNet)
	if !reflect.DeepEqual(ones, 24) {
		t.Fail()
	}
}

func TestAdd01(t *testing.T) {
	s1 := []byte{123, 0, 0, 1}
	s2 := []byte{123, 0, 0, 2}
	if !reflect.DeepEqual(add(s1, 1), s2) {
		t.Fail()
	}
}

func TestAdd02(t *testing.T) {
	s1 := []byte{123, 0, 0, 254}
	s2 := []byte{123, 0, 0, 255}
	if !reflect.DeepEqual(add(s1, 1), s2) {
		t.Fail()
	}
}

func TestAdd03(t *testing.T) {
	s1 := []byte{123, 0, 0, 255}
	s2 := []byte{123, 0, 1, 0}
	if !reflect.DeepEqual(add(s1, 1), s2) {
		t.Fail()
	}
}

func TestAdd04(t *testing.T) {
	s1 := []byte{123, 255, 255, 255}
	s2 := []byte{124, 0, 0, 0}
	if !reflect.DeepEqual(add(s1, 1), s2) {
		t.Fail()
	}
}

func TestAdd05(t *testing.T) {
	s1 := []byte{255, 255, 255, 255}
	s2 := []byte{0, 0, 0, 0}
	if !reflect.DeepEqual(add(s1, 1), s2) {
		t.Fail()
	}
}

func TestDec01(t *testing.T) {
	s1 := []byte{123, 0, 0, 1}
	s2 := []byte{123, 0, 0, 0}
	if !reflect.DeepEqual(dec(s1, 1), s2) {
		t.Fail()
	}
}

func TestDec02(t *testing.T) {
	s1 := []byte{123, 0, 1, 0}
	s2 := []byte{123, 0, 0, 255}
	if !reflect.DeepEqual(dec(s1, 1), s2) {
		t.Fail()
	}
}

func TestDec03(t *testing.T) {
	s1 := []byte{1, 0, 0, 0}
	s2 := []byte{0, 255, 255, 255}
	if !reflect.DeepEqual(dec(s1, 1), s2) {
		t.Fail()
	}
}

func TestDec04(t *testing.T) {
	s1 := []byte{3, 0, 0, 0}
	s2 := []byte{2, 255, 255, 255}
	if !reflect.DeepEqual(dec(s1, 1), s2) {
		t.Fail()
	}
}

func TestDec05(t *testing.T) {
	s1 := []byte{0, 0, 0, 0}
	s2 := []byte{255, 255, 255, 255}
	if !reflect.DeepEqual(dec(s1, 1), s2) {
		t.Fail()
	}
}

func TestGetBetweenRandomIp01(t *testing.T) {
	x1 := []byte{10, 1, 1, 2}
	x2 := []byte{10, 1, 1, 2}
	x, _ := getBetweenRandomIp(x1, x2)
	if !reflect.DeepEqual(x, x1) {
		t.Fail()
	}
}

func TestGetBetweenRandomIp02(t *testing.T) {
	x1 := []byte{10, 1, 2, 3}
	x2 := []byte{10, 1, 2}
	_, err := getBetweenRandomIp(x1, x2)
	if err == nil {
		t.Fail()
	}
}

func TestGetBetweenRandomIp03(t *testing.T) {
	x1 := []byte{10, 1, 2, 3}
	x2 := []byte{10, 1, 2, 78}
	x, err := getBetweenRandomIp(x1, x2)
	if err != nil {
		t.Fail()
	}
	if !(bytes.Compare(x1, x) <= 0 && bytes.Compare(x2, x) >= 0) {
		t.Fail()
	}
}

func TestGetBetweenRandomIp04(t *testing.T) {
	x1 := []byte{10, 1, 2, 3}
	x2 := []byte{10, 1, 20, 78}
	x, err := getBetweenRandomIp(x1, x2)
	if err != nil {
		t.Fail()
	}
	if !(bytes.Compare(x1, x) <= 0 && bytes.Compare(x2, x) >= 0) {
		t.Fail()
	}
}
