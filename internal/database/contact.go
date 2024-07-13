package database

import "portfolio/internal/types"

func (d *service) AddContact(req types.AddContactReq) error {
	contact := Contact{
		Name:    req.Name,
		Email:   req.Email,
		Message: req.Message,
	}

	if err := d.db.Create(&contact).Error; err != nil {
		return err
	}

	return nil
}

func (d *service) GetContacts() ([]Contact, error) {
	var contacts []Contact
	if err := d.db.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}
