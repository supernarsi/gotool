package encrypt

import (
	"crypto/sha256"

	"github.com/google/uuid"
)

func StringToUuid(inputStr string) string {
	hash := sha256.New()
	hash.Write([]byte(inputStr))
	hashBytes := hash.Sum(nil)
	u, err := uuid.FromBytes(hashBytes[:16])
	if err != nil {
		return ""
	}
	return u.String()
}

func StringToUuidByNamespace(namespaceUUID uuid.UUID, s string) string {
	// 使用命名空间 UUID 和字符串生成 UUIDv5
	return uuid.NewSHA1(namespaceUUID, []byte(s)).String()
}
