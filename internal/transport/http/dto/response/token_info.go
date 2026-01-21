package response

type TokenInfo struct {
	Token     string `json:"token"`
	Type      string `json:"type"`
	ExpiresIn int    `json:"expires_in"`
}
