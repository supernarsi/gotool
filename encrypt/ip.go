package encrypt

import (
	"errors"
	"math"
	"net"
	"strconv"
)

var ipEncryptKey uint64 = 42

func IpEncrypt(ip string) string {
	ipInt, err := IPString2Long(ip)
	if err != nil {
		return ""
	}
	return strconv.FormatUint(uint64(ipInt)+ipEncryptKey, 16)
}

func IpDecrypt(ipStr string) string {
	ipInt, err := strconv.ParseUint(ipStr, 16, 64)

	if err != nil {
		return ""
	}
	ip, err := Long2IPString(uint(ipInt - ipEncryptKey))
	if err != nil {
		return ""
	}
	return ip
}

func IPString2Long(ip string) (uint, error) {
	b := net.ParseIP(ip).To4()
	if b == nil {
		return 0, errors.New("beyond the scope of ipv4")
	}

	return uint(b[3]) | uint(b[2])<<8 | uint(b[1])<<16 | uint(b[0])<<24, nil
}

func Long2IPString(i uint) (string, error) {
	if i > math.MaxUint32 {
		return "", errors.New("beyond the scope of ipv4")
	}

	ip := make(net.IP, net.IPv4len)
	ip[0] = byte(i >> 24)
	ip[1] = byte(i >> 16)
	ip[2] = byte(i >> 8)
	ip[3] = byte(i)

	return ip.String(), nil
}
