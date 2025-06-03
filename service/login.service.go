package service

import (
	"api-crud-golang/model"
	"api-crud-golang/util"
	"log"
)

func (s *MainService) LoginService(email, password string) (string, error) {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println(err)
		return "", err
	}

	if !util.CheckPassword(password, user.Password) {
		log.Println("Invalid email or password")
		return "", nil
	}

	token, err := util.GenreateJWT(user.Email, user.Password)
	if err != nil {
		log.Println(err)
		return "", nil
	}

	return token, nil
}
