package services

import (
	"errors"
	"tuxedo/database"
	"tuxedo/models/entity"

	"gorm.io/gorm"
)

func GetUserByID(id uint) (*entity.Users, error) {
	var user entity.Users
	err := database.DB.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
