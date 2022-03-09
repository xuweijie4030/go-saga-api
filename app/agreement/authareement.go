package agreement

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type RefreshRequest struct {
	Token string `json:"token"`
}

type RefreshResponse struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}
