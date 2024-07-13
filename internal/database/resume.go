package database

import "portfolio/internal/types"

func (d *service) AddResume(req types.AddResumeReq) error {
	resume := Resume{
		Title: req.Title,
		Link:  req.Link,
	}

	if err := d.db.Create(&resume).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) DeleteResume(id uint) error {
	if err := d.db.Delete(&Resume{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) GetPublishedResume() (Resume, error) {
	if err := d.db.Where("published = ?", true).First(&Resume{}).Error; err != nil {
		return Resume{}, err
	}
	return Resume{}, nil
}

func (d *service) GetAllResume() ([]Resume, error) {
	var resumes []Resume
	if err := d.db.Find(&resumes).Error; err != nil {
		return nil, err
	}
	return resumes, nil
}
