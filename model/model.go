package model

import (
	"time"
)

type SavedFile struct {
	Filename               string
	InitTime               time.Time
	ExpireDuration         time.Duration
	AccessedExpireDuration time.Duration
}
