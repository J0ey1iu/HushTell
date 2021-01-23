package model

import (
	"time"
)

// CachedInfo is the info we store in cache
// swagger:model CachedInfo
type CachedInfo struct {
	// the original filename of the uploaded file
	Filename string `json:"filename"`

	// the generated url for the file
	Url string `json:"Url"`

	// the extension of the file, we wouldn't want to show this to users
	Extension string `json:"ext"`

	// the upload time of the file
	UploadTime time.Time `json:"upload-time"`

	// how long do we keep this file after being accessed
	Duration time.Duration `json:"duration"`

	// user's input when upload this file
	Options Option `json:"-"`
}

// CachedFiles our main cache
var CachedFiles map[string]CachedInfo = make(map[string]CachedInfo)
