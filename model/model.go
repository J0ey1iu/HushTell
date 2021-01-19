package model

import (
	"time"
)

type SavedFile struct {
	Filename               string
	Url                    string
	Extension              string
	InitTime               time.Time
	ExpireDuration         time.Duration
	AccessedExpireDuration time.Duration
}

type Test struct {
	IP       string
	HashedIP string
}
