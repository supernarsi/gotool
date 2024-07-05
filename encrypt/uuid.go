package encrypt

import (
	"crypto/sha256"

	"github.com/google/uuid"
)

func StringToUUID(inputStr string) string {
	hash := sha256.New()
	hash.Write([]byte(inputStr))
	hashBytes := hash.Sum(nil)
	u, err := uuid.FromBytes(hashBytes[:16])
	if err != nil {
		return ""
	}
	return u.String()
}
