package model

import (
	"mime/multipart"
	"time"
)

// TextUploadFormat defines the not upload format for api documentation
// swagger:parameters uploadText
type TextUploadFormat struct {
	// in:formData
	// the note user typed in
	Text string `json:"mytext"`

	// in:formData
	// the options user specified, submitted as stringified JSON format
	Options string `json:"options"`
}

// Option is the user options we provides
// swagger:model Option
type Option struct {
	// in:query
	EmailTip      bool          `json:"emailTip"`
	ReadTip       bool          `json:"readTip"`
	Encryption    bool          `json:"encryption"`
	EncryptionPWD string        `json:"pwd"`
	ReadTime      time.Duration `json:"readTime"`
	SaveTime      time.Duration `json:"saveTime"`
}

// FileUploadFormat defines the upload format for api documentation
// swagger:parameters uploadFile
type FileUploadFormat struct {
	// in:formData
	// swagger:file myfile
	File *multipart.FileHeader `json:"myfile"`

	// in:formData
	// the options user specified, submitted as stringified JSON format
	Options string `json:"options"`
}
