package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateTokens(email string, role string) (*TokenResponse, error)
	RefreshToken(refreshToken string) (*TokenResponse, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
	Role         string `json:"role"`
}

type jwtClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey       string
	refreshKey      string
	issuer          string
	accessTokenExp  time.Duration
	refreshTokenExp time.Duration
}

// -------------------------------------------------------------
// 1) GÉNÉRATION DES TOKENS : ACCESS + REFRESH
// -------------------------------------------------------------
func (j *jwtService) GenerateTokens(email string, role string) (*TokenResponse, error) {

	// Access token (1h)
	accessExp := time.Now().Add(j.accessTokenExp).Unix()

	accessClaims := &jwtClaims{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExp,
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessString, err := accessToken.SignedString([]byte(j.secretKey))
	if err != nil {
		return nil, err
	}

	// Refresh token (7 jours)
	refreshExp := time.Now().Add(j.refreshTokenExp).Unix()

	refreshClaims := jwt.StandardClaims{
		ExpiresAt: refreshExp,
		Issuer:    j.issuer,
		IssuedAt:  time.Now().Unix(),
		Subject:   email,
	}

	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshString, err := refreshTokenObj.SignedString([]byte(j.refreshKey))
	if err != nil {
		return nil, err
	}

	// Réponse
	return &TokenResponse{
		AccessToken:  accessString,
		RefreshToken: refreshString,
		ExpiresAt:    accessExp,
		Role:         role,
	}, nil
}

// -------------------------------------------------------------
// 2) UTILISER LE REFRESH TOKEN POUR GÉNÉRER UN NOUVEAU ACCESS TOKEN
// -------------------------------------------------------------
func (j *jwtService) RefreshToken(refreshTokenString string) (*TokenResponse, error) {

	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.refreshKey), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("refresh token invalide")
	}

	claims := token.Claims.(jwt.MapClaims)

	email := claims["sub"].(string)
	fmt.Println("Refresh for user:", email)

	// ⚠️ Ici tu récupères le role depuis ta base
	// => getUserRole(email)
	//admin := true // exemple
	return j.GenerateTokens(email, "admin")
}

// -------------------------------------------------------------
// 3) VALIDATION DU ACCESS TOKEN
// -------------------------------------------------------------
func (j *jwtService) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})
}

// -------------------------------------------------------------
// INIT DU SERVICE
// -------------------------------------------------------------
func NewJWTService() JWTService {
	return &jwtService{
		secretKey:       getEnv("JWT_SECRET", "access_secret"),
		refreshKey:      getEnv("JWT_REFRESH_SECRET", "refresh_secret"),
		issuer:          time.Hour.String(),
		accessTokenExp:  15 * time.Minute,   // 1h
		refreshTokenExp: 24 * 7 * time.Hour, // 7 jours

	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
