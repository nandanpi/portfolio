package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) AddWork(req types.AddWorkReq) error {
	// Create the Work record
	work := Work{
		Title:       req.Title,
		Description: req.Description,
		GithubLink:  req.GithubLink,
		DemoLink:    req.DemoLink,
		Image:       req.Image,
		Published:   false, // or use req.Published if you want to handle this dynamically
	}

	tx := d.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&work).Error; err != nil {
		tx.Rollback()
		return err
	}

	var techs []*Technology
	if err := tx.Where("id IN ?", req.TechStack).Find(&techs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&work).Association("TechStack").Replace(techs); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (d *service) UpdateWork(req types.UpdateWorkReq) error {
	// Retrieve the existing Work record
	var work Work
	if err := d.db.First(&work, req.ID).Error; err != nil {
		return err
	}

	work.Title = req.Title
	work.Description = req.Description
	work.GithubLink = req.GithubLink
	work.DemoLink = req.DemoLink
	work.Image = req.Image

	tx := d.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&work).Error; err != nil {
		tx.Rollback()
		return err
	}

	var techs []*Technology
	if err := tx.Where("id IN ?", req.TechStack).Find(&techs).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&work).Association("TechStack").Replace(techs); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}

func (d *service) DeleteWork(id uint) error {
	if err := d.db.Delete(&Work{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) ToggleWorkPublished(id uint) error {
	work := Work{
		Model: gorm.Model{
			ID: id,
		},
	}

	if err := d.db.Model(&work).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) GetPublishedWorks() ([]Work, error) {
	var works []Work

	if err := d.db.Where("published = ?", true).Preload("TechStack").Find(&works).Error; err != nil {
		return nil, err
	}

	return works, nil
}

func (d *service) GetAllWorks() ([]Work, error) {
	var works []Work

	if err := d.db.Preload("TechStack").Find(&works).Error; err != nil {
		return nil, err
	}

	return works, nil
}
