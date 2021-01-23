package model

// AccessResponse for file access, tells whether the access was successful or not
// swagger:response AccessResult
type AccessResponse struct {
	Result string `json:"result"`
}
