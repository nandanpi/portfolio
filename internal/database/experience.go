package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) AddExperience(req types.AddExperienceReq) error {
	experience := Experience{
		Title:       req.Title,
		Description: req.Description,
		Company:     req.Company,
		Location:    req.Location,
		Logo:        req.Logo,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}
	if err := d.db.Create(&experience).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) UpdateExperience(req types.UpdateExperienceReq) error {
	experience := Experience{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:       req.Title,
		Description: req.Description,
		Company:     req.Company,
		Location:    req.Location,
		Logo:        req.Logo,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
	}
	if err := d.db.Save(&experience).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) DeleteExperience(id uint) error {
	if err := d.db.Delete(&Experience{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) ToggleExperiencePublish(id uint) error {
	if err := d.db.Model(&Experience{}).Where("id = ?", id).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) GetPublishedExperiences() ([]Experience, error) {
	var experiences []Experience
	if err := d.db.Where("published = ?", true).Find(&experiences).Error; err != nil {
		return nil, err
	}
	return experiences, nil
}

func (d *service) GetAllExperiences() ([]Experience, error) {
	var experiences []Experience
	if err := d.db.Find(&experiences).Error; err != nil {
		return nil, err
	}
	return experiences, nil
}
