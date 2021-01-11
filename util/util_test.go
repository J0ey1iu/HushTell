package util

import (
	"fmt"
	"testing"
)

func TestEncodeIP(t *testing.T) {
	hashed := encodeIP("192.168.0.1")
	fmt.Println(hashed)
}
