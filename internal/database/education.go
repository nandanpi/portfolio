package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) CreateEducation(req types.AddEducationReq) error {
	education := Education{
		School:    req.School,
		Degree:    req.Degree,
		Aggregate: req.Aggregate,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
	if err := d.db.Create(&education).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) UpdateEducation(req types.UpdateEducationReq) error {
	education := Education{
		Model: gorm.Model{
			ID: req.ID,
		},
		School:    req.School,
		Degree:    req.Degree,
		Aggregate: req.Aggregate,
		StartDate: req.StartDate,
		EndDate:   req.EndDate,
	}
	if err := d.db.Save(&education).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) DeleteEducation(id uint) error {
	if err := d.db.Delete(&Education{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) ToggleEducationPublished(id uint) error {
	if err := d.db.Model(&Education{}).Where("id = ?", id).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}
	return nil
}

func (d *service) GetPublishedEducations() ([]Education, error) {
	var educations []Education

	if err := d.db.Where("published = ?", true).Find(&educations).Error; err != nil {
		return nil, err
	}

	return educations, nil
}

func (d *service) GetAllEducations() ([]Education, error) {
	var educations []Education

	if err := d.db.Find(&educations).Error; err != nil {
		return nil, err
	}

	return educations, nil
}
