package user

import (
	"gin_web_api/utils"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	FindByID(ID int) (User, error)
	Login(loginrequest LoginRequest) (string, error)
	Create(registerrequest RegisterRequest) (User, error)
	// Update(ID int, registerrequest RegisterRequest) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByID(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	return user, err
}

func (s *service) Create(registerrequest RegisterRequest) (User, error) {

	newUser, _ := hashedUser(User{
		Username: registerrequest.Username,
		Password: registerrequest.Password,
	})

	user, err := s.repository.Create(newUser)

	return user, err
}

func (s *service) Login(loginrequest LoginRequest) (string, error) {

	userCheck, err := s.repository.FindByUsername(loginrequest.Username)
	if err != nil {
		return "", err
	}

	err = verifyPassword(loginrequest.Password, userCheck.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(userCheck.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func hashedUser(u User) (User, error) {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return u, err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return u, nil

}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
