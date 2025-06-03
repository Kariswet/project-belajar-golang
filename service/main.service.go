package service

import (
	"api-crud-golang/model"
	"api-crud-golang/util"
	"log"

	"gorm.io/gorm"
)

type MainService struct {
	db *gorm.DB
}

func NewMainService(db *gorm.DB) *MainService {
	return &MainService{
		db: db,
	}
}

func (s *MainService) ServicePostMainModel(data *model.MainModel) error {
	data.Id = util.GenerateId()

	if err := s.db.Create(data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *MainService) ServiceGetMainModel() ([]model.MainModel, error) {
	var main []model.MainModel

	if err := s.db.Find(&main).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return main, nil
}

func (s *MainService) ServiceGetByIdMainModel(id string) (*model.MainModel, error) {
	var main model.MainModel

	if err := s.db.First(&main, "id = ?", id).Error; err != nil {
		log.Println(err)
		return nil, err
	}

	return &main, nil
}

func (s *MainService) ServiceUpdateMainModel(id string, data *model.MainModel) error {
	var update model.MainModel

	if err := s.db.First(&update, "id = ?", id).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := s.db.Model(&update).Updates(&data).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *MainService) ServiceDeleteMainModel(id string) error {
	var main model.MainModel

	if err := s.db.First(&main, "id = ?", id).Error; err != nil {
		log.Println(err)
		return err
	}

	if err := s.db.Delete(&main).Error; err != nil {
		log.Println(err)
		return err
	}

	return nil
}
