package service

import "errors"

type AuthService struct {
    // В учебной версии — фиксированный токен
    validUsername string
    validPassword string
    token         string
}

func NewAuthService() *AuthService {
    return &AuthService{
        validUsername: "student",
        validPassword: "student",
        token:         "demo-token",
    }
}

type LoginRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type LoginResponse struct {
    AccessToken string `json:"access_token"`
    TokenType   string `json:"token_type"`
}

type VerifyResponse struct {
    Valid   bool   `json:"valid"`
    Subject string `json:"subject,omitempty"`
    Error   string `json:"error,omitempty"`
}

func (s *AuthService) Login(req LoginRequest) (LoginResponse, error) {
    if req.Username != s.validUsername || req.Password != s.validPassword {
        return LoginResponse{}, errors.New("invalid credentials")
    }
    return LoginResponse{
        AccessToken: s.token,
        TokenType:   "Bearer",
    }, nil
}

func (s *AuthService) VerifyToken(token string) (VerifyResponse, error) {
    if token == "" {
        return VerifyResponse{
            Valid: false,
            Error: "unauthorized",
        }, nil
    }
    if token != s.token {
        return VerifyResponse{
            Valid: false,
            Error: "unauthorized",
        }, nil
    }
    return VerifyResponse{
        Valid:   true,
        Subject: "student",
    }, nil
}
