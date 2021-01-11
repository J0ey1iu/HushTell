package util

import (
	"encoding/base64"
)

func encodeIP(address string) string {
	return base64.StdEncoding.EncodeToString([]byte(address))
}
