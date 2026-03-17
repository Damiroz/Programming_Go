package service

type AuthService struct{}

func NewAuthService() *AuthService {
    return &AuthService{}
}

func (s *AuthService) Login(username, password string) (string, bool) {
    if username == "admin" && password == "admin" {
        return "demo-token", true
    }
    return "", false
}

func (s *AuthService) ValidateToken(token string) (bool, string) {
    if token == "demo-token" {
        return true, "student"
    }
    return false, ""
}
