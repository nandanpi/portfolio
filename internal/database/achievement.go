package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) AddAchievement(req types.AddAchievementReq) error {
	achievement := Achievement{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Date:        req.Date,
	}
	if err := d.db.Create(&achievement).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) UpdateAchievement(req types.UpdateAchievementReq) error {
	achievement := Achievement{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Date:        req.Date,
	}
	if err := d.db.Save(&achievement).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) DeleteAchievement(id uint) error {
	if err := d.db.Delete(&Achievement{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) ToggleAchievementPublished(id uint) error {
	if err := d.db.Model(&Achievement{}).Where("id = ?", id).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) GetPublishedAchievements() ([]Achievement, error) {
	var achievements []Achievement

	if err := d.db.Where("published = ?", true).Find(&achievements).Error; err != nil {
		return nil, err
	}

	return achievements, nil
}

func (d *service) GetAllAchievements() ([]Achievement, error) {
	var achievements []Achievement

	if err := d.db.Find(&achievements).Error; err != nil {
		return nil, err
	}

	return achievements, nil
}
