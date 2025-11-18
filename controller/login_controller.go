package controller

import (
	"loyalty-api/controller/dto"
	service "loyalty-api/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

// ------------------------------------------------------------
// LOGIN : Vérifie email/password -> Génère Access+Refresh Tokens
// ------------------------------------------------------------
func (lc *loginController) Login(ctx *gin.Context) {

	var credentials dto.Credentials

	// 🔹 Validate JSON Body
	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// 🔹 Vérification du user
	user, ok := lc.loginService.GetUserByLogin(credentials.Email, credentials.Password)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// user contient maintenant Email + Role
	tokens, err := lc.jwtService.GenerateTokens(user.Email, user.Role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// 🔹 Envoi final
	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
		"expires_at":    tokens.ExpiresAt,
		"role":          tokens.Role,
	})
}

// ------------------------------------------------------------
// Constructeur
// ------------------------------------------------------------
func NewLoginController(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}
