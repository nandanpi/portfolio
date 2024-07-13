package database

func (d *service) GetUser(username string) (Users, error) {
	user := &Users{}
	result := d.db.Model(&Users{}).Where("username = ?", username).First(user)
	if result.Error != nil {
		return Users{}, result.Error
	}

	return *user, nil
}
