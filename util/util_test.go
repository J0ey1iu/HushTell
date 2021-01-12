package util

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hashed := Hash("192.168.0.1")
	fmt.Println(hashed)
}

func TestShortHash(t *testing.T) {
	hashed := ShortHash("192.168.0.1")
	fmt.Println(hashed)
}

func TestCreateFolderByName(t *testing.T) {
	CreateFolderByName("test")
}

func TestCreateFile(t *testing.T) {
	createFile()
}
