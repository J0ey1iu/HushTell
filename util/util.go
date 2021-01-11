package util

import (
	"crypto/sha1"
	"fmt"
	"os"
	"strings"
)

// Hash uses sha-1 to hash the input string
func Hash(s string) string {
	hasher := sha1.New()
	hasher.Write([]byte(s))
	bs := hasher.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

// ShortHash picks six characters from sha-1 hash
func ShortHash(s string) string {
	h := Hash(s)
	choices := []int{1, 3, 20, 21, 30, 39}
	ret := ""
	for _, c := range choices {
		ret += string(h[c])
	}
	return strings.ToUpper(ret)
}

// CreateFolderByName just as the name says
func CreateFolderByName(name string) {
	// first determine if the folder exists
	if _, err := os.Stat("./temp/" + name); os.IsNotExist(err) {
		os.MkdirAll("./temp/"+name, os.ModePerm)
	}
}
