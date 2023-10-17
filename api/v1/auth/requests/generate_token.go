package requests

import "chat-hex/business/auth"

type GenerateTokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req *GenerateTokenRequest) ToGenerateTokenSpec() *auth.GenerateTokenSpec {
	var generateTokenSpec auth.GenerateTokenSpec

	generateTokenSpec.Email = req.Email
	generateTokenSpec.Password = req.Password

	return &generateTokenSpec
}