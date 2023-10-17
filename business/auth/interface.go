package auth

type Service interface {
	GenerateJWT(generateTokenSpec GenerateTokenSpec) (string, error)
	ValidateToken(signedToken string) error
}