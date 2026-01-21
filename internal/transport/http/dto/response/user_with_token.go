package response

type ResultResponseWithToken struct {
	User  *ReadUserResponse `json:"user"`
	Token *TokenInfo        `json:"token_info,omitempty"`
}
