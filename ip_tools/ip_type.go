package ip_tools

type IpType int32

const (
	UNKNOWN IpType = 0
	IPV4    IpType = 1
	IPV6    IpType = 2
)

var IpType_names = map[IpType]string{
	UNKNOWN: "unknown_ip_type",
	IPV4:    "ipv4",
	IPV6:    "ipv6",
}

var IpType_values = map[string]IpType{
	"unknown_ip_type": UNKNOWN,
	"ipv4":            IPV4,
	"ipv6":            IPV6,
}

func (i IpType) String() string {
	return IpType_names[i]
}
