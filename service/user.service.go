package service

import (
	"api-crud-golang/model"
	"api-crud-golang/util"
	"log"
)

func (s *MainService) ServicePostUser(data *model.User) error {
	hashPwd, err := util.HashPassword(data.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	data.Password = hashPwd
	data.ID = util.GenerateId()

	if err := s.db.Create(data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *MainService) ServiceDeleteUser(id string) error {
	var model model.User

	if err := s.db.First(&model, "id = ?", id).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := s.db.Delete(&model).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
