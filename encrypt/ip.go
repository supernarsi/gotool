package encrypt

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net"
)

const (
	ipv4len         = net.IPv4len
	ipEncryptKeyLen = 16
	ipEncryptKeyDef = "rNd3TQq9BdTAKF42"
)

var ErrIPBeyondIPv4 = errors.New("IP address beyond the scope of IPv4")
var ErrInvalidIPHex = errors.New("invalid IP address hex")

func newIPKey() ([]byte, error) {
	key := make([]byte, ipEncryptKeyLen)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

type IPEncrypter struct {
	key []byte
}

func NewIPEncrypter() *IPEncrypter {
	return &IPEncrypter{key: []byte(ipEncryptKeyDef)}
}

func (enc *IPEncrypter) Encrypt(ip string) (string, error) {
	ipBytes := net.ParseIP(ip).To4()
	if ipBytes == nil {
		return "", ErrIPBeyondIPv4
	}
	ipUint32 := uint32(ipBytes[0])<<24 | uint32(ipBytes[1])<<16 | uint32(ipBytes[2])<<8 | uint32(ipBytes[3])
	ipUint64 := uint64(ipUint32) + enc.ipEncryptKey()
	ipHex := hex.EncodeToString(enc.encodeIPUint64(ipUint64))
	return ipHex, nil
}

func (enc *IPEncrypter) Decrypt(ipHex string) (string, error) {
	if ipHex == "" {
		return "", nil
	}
	ipBytes, err := hex.DecodeString(ipHex)
	if err != nil {
		return "", ErrInvalidIPHex
	}
	ipUint64 := enc.decodeIPUint64(ipBytes)
	ipUint32 := uint32(ipUint64 - enc.ipEncryptKey())
	ip := net.IPv4(byte(ipUint32>>24), byte(ipUint32>>16), byte(ipUint32>>8), byte(ipUint32))
	return ip.String(), nil
}

func (enc *IPEncrypter) ipEncryptKey() uint64 {
	return enc.decodeIPUint64(enc.key)
}

func (enc *IPEncrypter) encodeIPUint64(ipUint64 uint64) []byte {
	b := make([]byte, 8)
	b[0] = byte(ipUint64 >> 56)
	b[1] = byte(ipUint64 >> 48)
	b[2] = byte(ipUint64 >> 40)
	b[3] = byte(ipUint64 >> 32)
	b[4] = byte(ipUint64 >> 24)
	b[5] = byte(ipUint64 >> 16)
	b[6] = byte(ipUint64 >> 8)
	b[7] = byte(ipUint64)
	return b
}

func (enc *IPEncrypter) decodeIPUint64(b []byte) uint64 {
	return (uint64(b[0]) << 56) | (uint64(b[1]) << 48) | (uint64(b[2]) << 40) | (uint64(b[3]) << 32) | (uint64(b[4]) << 24) | (uint64(b[5]) << 16) | (uint64(b[6]) << 8) | uint64(b[7])
}
