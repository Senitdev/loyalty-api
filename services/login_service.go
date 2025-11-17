package service

import "loyalty-api/repository"

type LoginService interface {
	GetUserByLogin(username string, password string) bool
}
type loginService struct {
	service repository.LoginRepository
}

// GetUserByLoginAndPassword implements Loginv2Service.
func (l *loginService) GetUserByLogin(username string, password string) bool {
	result := l.service.GetUserByLoginAndPassword(username, password)
	if !result {
		return false
	}
	return true

}

// Login implements Loginv2Service.
func NewLoginService(repo repository.LoginRepository) LoginService {
	return &loginService{
		service: repo,
	}
}
