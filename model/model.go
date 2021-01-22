package model

import (
	"time"
)

type CachedInfo struct {
	Filename   string        `json:"filename"`
	Url        string        `json:"Url"`
	Extension  string        `json:"ext"`
	UploadTime time.Time     `json:"upload-time"`
	Duration   time.Duration `json:"duration"`
}

type Test struct {
	IP       string
	HashedIP string
}
