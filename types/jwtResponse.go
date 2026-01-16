package types

type JwtResponse struct {
	Jwt          string `json:"jwt,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
