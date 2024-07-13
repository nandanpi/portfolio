package database

import (
	"portfolio/internal/types"

	"gorm.io/gorm"
)

func (d *service) AddExtraCurricular(req types.AddExtraCurricularReq) error {
	extraCurricular := ExtraCurricular{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Date:        req.Date,
	}

	if err := d.db.Create(&extraCurricular).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) UpdateExtraCurricular(req types.UpdateExtraCurricularReq) error {
	extraCurricular := ExtraCurricular{
		Model: gorm.Model{
			ID: req.ID,
		},
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Date:        req.Date,
	}

	if err := d.db.Save(&extraCurricular).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) DeleteExtraCurricular(id uint) error {
	if err := d.db.Delete(&ExtraCurricular{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) ToggleExtraCurricularPublish(id uint) error {
	if err := d.db.Model(&ExtraCurricular{}).Where("id = ?", id).Update("published", gorm.Expr("NOT published")).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) GetPublishedExtraCurriculars() ([]ExtraCurricular, error) {
	var extraCurriculars []ExtraCurricular

	if err := d.db.Where("published = ?", true).Find(&extraCurriculars).Error; err != nil {
		return nil, err
	}

	return extraCurriculars, nil
}

func (d *service) GetAllExtraCurriculars() ([]ExtraCurricular, error) {
	var extraCurriculars []ExtraCurricular

	if err := d.db.Find(&extraCurriculars).Error; err != nil {
		return nil, err
	}

	return extraCurriculars, nil
}
