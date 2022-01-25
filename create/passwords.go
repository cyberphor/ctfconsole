package create

import (
	"crypto/md5"
	"encoding/hex"
)

func HashPassword(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
