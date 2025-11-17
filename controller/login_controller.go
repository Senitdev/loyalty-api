package controller

import (
	"loyalty-api/controller/dto"
	service "loyalty-api/services"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}
type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

// Login implements LoginController.
func (l *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return err.Error()
	}
	result := l.loginService.GetUserByLogin(credentials.Username, credentials.Password)
	if !result {
		return ""
	}
	return l.jWtService.GenerateToken(credentials.Username, true)
}

func NewLoginController(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}
