package ip_tools

import (
	"fmt"
	"testing"
)

var (
	CheckIpErrorFormat     = "ip: '%s' format is incorrect"
	CheckIpCidrErrorFormat = "ip cidr: '%s' format is incorrect"
)

func TestCheckValidIp01(t *testing.T) {
	x := "127.0.0.1"
	err := CheckValidIp(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIp02(t *testing.T) {
	x := "1270.0.0.1"
	err := CheckValidIp(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIp03(t *testing.T) {
	x := "x.8.9.0"
	err := CheckValidIp(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIp04(t *testing.T) {
	x := "0.0.0.0"
	err := CheckValidIp(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIp05(t *testing.T) {
	x := "255.255.255.255"
	err := CheckValidIp(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIp06(t *testing.T) {
	x := "255.255.255"
	err := CheckValidIp(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIp07(t *testing.T) {
	x := "ABCD:EF01:2345:6789:ABCD:EF01:2345:6789"
	err := CheckValidIp(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIpCidr01(t *testing.T) {
	x := "127.0.0.2/24"
	err := CheckValidIpCidr(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIpCidr02(t *testing.T) {
	x := "127.0.0.200/24"
	err := CheckValidIpCidr(x)
	if err != nil {
		t.Fail()
	}
}

func TestCheckValidIpCidr03(t *testing.T) {
	x := "127.0.0.2"
	err := CheckValidIpCidr(x)
	// fmt.Println(err)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpCidrErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIpCidr04(t *testing.T) {
	x := "127.0.0.200/33"
	err := CheckValidIpCidr(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpCidrErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIpCidr05(t *testing.T) {
	x := "127.0.0/33"
	err := CheckValidIpCidr(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpCidrErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIpCidr06(t *testing.T) {
	x := "ABCD:EF01:2345:6789:ABCD:EF01:2345:6789/129"
	err := CheckValidIpCidr(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpCidrErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIpCidr07(t *testing.T) {
	x := "ABCD:EF01:2345:6789:ABCD:EF01:2345/34"
	err := CheckValidIpCidr(x)
	if err == nil {
		t.Fail()
	} else if err.Error() != fmt.Sprintf(CheckIpCidrErrorFormat, x) {
		t.Fail()
	}
}

func TestCheckValidIpCidr08(t *testing.T) {
	x := "ABCD:EF01:2345:6789:ABCD:EF01:2345:122/34"
	err := CheckValidIpCidr(x)
	if err != nil {
		t.Fail()
	}
}

func TestGetValidIpType01(t *testing.T) {
	if GetValidIpType("1.1.1.1") != IPV4 {
		t.Fail()
	}
}

func TestGetValidIpType02(t *testing.T) {
	if GetValidIpType("1:1:1:1:1:1:1:1") != IPV6 {
		t.Fail()
	}
}

func TestGetValidIpCidrType01(t *testing.T) {
	if GetValidIpType("1.1.1.1/24") != IPV4 {
		t.Fail()
	}
}

func TestGetValidIpCidrType02(t *testing.T) {
	if GetValidIpType("1:1:1:1:1:1:1:1/24") != IPV6 {
		t.Fail()
	}
}

func TestGetRandomIpV4(t *testing.T) {
	x := GetRandomIpV4()
	if _, err := parseIp(x.String()); err != nil {
		t.Fail()
	}

}

func TestGetRandomIpV6(t *testing.T) {
	x := GetRandomIpV6()
	if _, err := parseIp(x.String()); err != nil {
		t.Fail()
	}
}

func TestGetRandomIpCidrV4(t *testing.T) {
	x := GetRandomIpCidrV4()
	if _, err := parseIpCidr(x.String()); err != nil {
		t.Fail()
	}

}

func TestGetRandomIpCidrV6(t *testing.T) {
	x := GetRandomIpCidrV6()
	if _, err := parseIpCidr(x.String()); err != nil {
		t.Fail()
	}
}
