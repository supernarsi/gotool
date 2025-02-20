package encrypt

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"net"
)

const (
	ipv4len         = net.IPv4len
	ipEncryptKeyLen = 16
)

var (
	ErrIPBeyondIPv4 = errors.New("IP address beyond the scope of IPv4")
	ErrInvalidIPHex = errors.New("invalid IP address hex")
	defaultIPKey    = []byte{0x72, 0x4E, 0x64, 0x33, 0x54, 0x51, 0x71, 0x39, 0x42, 0x64, 0x54, 0x41, 0x4B, 0x46, 0x34, 0x32} // "rNd3TQq9BdTAKF42" 的 hex 表示
)

// 生成新的 16 字节密钥
func newIPKey() ([]byte, error) {
	key := make([]byte, ipEncryptKeyLen)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	return key, nil
}

type IPEncryptor struct {
	key uint64
}

// NewIPEncryptorWithKey 允许传入密钥，未提供则使用默认密钥
func NewIPEncryptorWithKey(hexKey string) (*IPEncryptor, error) {
	var keyBytes []byte
	var err error

	if hexKey == "" {
		keyBytes = defaultIPKey
	} else {
		keyBytes, err = hex.DecodeString(hexKey)
		if err != nil || len(keyBytes) != 8 {
			return nil, errors.New("invalid encryption key, must be 8 bytes hex")
		}
	}

	return &IPEncryptor{key: binary.BigEndian.Uint64(keyBytes)}, nil
}

// NewIPEncryptor 使用默认密钥
func NewIPEncryptor() *IPEncryptor {
	return &IPEncryptor{key: binary.BigEndian.Uint64(defaultIPKey)}
}

// Encrypt 加密 IP 地址
func (enc *IPEncryptor) Encrypt(ip string) (string, error) {
	ipBytes := net.ParseIP(ip).To4()
	if ipBytes == nil {
		return "", ErrIPBeyondIPv4
	}

	ipUint32 := binary.BigEndian.Uint32(ipBytes)
	ipUint64 := uint64(ipUint32) + enc.key

	// 使用 BigEndian 直接转换
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], ipUint64)
	return hex.EncodeToString(buf[:]), nil
}

// Decrypt 解密 IP 地址
func (enc *IPEncryptor) Decrypt(ipHex string) (string, error) {
	if ipHex == "" {
		return "", errors.New("empty encrypted IP string")
	}

	ipBytes, err := hex.DecodeString(ipHex)
	if err != nil || len(ipBytes) != 8 {
		return "", ErrInvalidIPHex
	}

	ipUint64 := binary.BigEndian.Uint64(ipBytes)
	if ipUint64 < enc.key {
		return "", errors.New("invalid encrypted IP value")
	}

	ipUint32 := uint32(ipUint64 - enc.key)

	// 还原 IP 地址
	ip := net.IPv4(byte(ipUint32>>24), byte(ipUint32>>16), byte(ipUint32>>8), byte(ipUint32))
	return ip.String(), nil
}
