package service

import (
	"loyalty-api/controller/dto"
	"loyalty-api/repository"
)

type LoginService interface {
	GetUserByLogin(email string, password string) (dto.UserDTO, bool)
}
type loginService struct {
	service repository.LoginRepository
}

// GetUserByLoginAndPassword implements Loginv2Service.
func (l *loginService) GetUserByLogin(email string, password string) (dto.UserDTO, bool) {
	result, ok := l.service.GetUserByLoginAndPassword(email, password)
	if !ok {
		return dto.UserDTO{}, false
	}
	return result, true
}

// Login implements Loginv2Service.
func NewLoginService(repo repository.LoginRepository) LoginService {
	return &loginService{
		service: repo,
	}
}
