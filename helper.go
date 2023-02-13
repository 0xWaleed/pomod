package pomod

import (
	"crypto/rand"
	"encoding/hex"
)

func generateId() string {
	length := 10
	buffer := make([]byte, length/2)
	_, err := rand.Read(buffer)
	if err != nil {
		return ""
	}
	out := hex.EncodeToString(buffer)
	return out
}
