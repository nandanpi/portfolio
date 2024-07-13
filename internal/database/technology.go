package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) AddTechnology(req types.AddTechnologyReq) error {
	technology := Technology{
		Name: req.Name,
		Logo: req.Logo,
	}

	if err := d.db.Create(&technology).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) DeleteTechnology(id uint) error {
	if err := d.db.Delete(&Technology{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) ToggleTechnologyStatus(id uint) error {
	if err := d.db.Model(&Technology{}).Where("id = ?", id).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) GetPublishedTechnologies() ([]Technology, error) {
	var technologies []Technology
	if err := d.db.Preload("Works").Where("published = ?", true).Find(&technologies).Error; err != nil {
		return nil, err
	}
	return technologies, nil
}

func (d *service) GetAllTechnologies() ([]Technology, error) {
	var technologies []Technology
	if err := d.db.Preload("Works").Find(&technologies).Error; err != nil {
		return nil, err
	}
	return technologies, nil
}
