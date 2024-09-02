package services

import (
	"tuxedo/database"
	"tuxedo/models/entity"
	"tuxedo/models/request"

	"gorm.io/gorm"
)

var DB *gorm.DB

func GetUserByID(userID uint) (*entity.Users, error) {
	var user entity.Users
	err := database.DB.Preload("Contacts").First(&user, userID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func BuildUserProfile(user *entity.Users) (request.UserProfile, error) {
	var contacts []request.Contact
	for _, contact := range user.Contacts {
		contacts = append(contacts, request.Contact{
			Phone: contact.Phone,
			Bio:   contact.Bio,
		})
	}

	profile := request.UserProfile{
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02"),
		Contacts:  contacts,
	}

	return profile, nil
}

func UpdateUserProfile(updateRequest request.UpdateUserProfileRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		user := entity.Users{
			ID:    updateRequest.ID,
			Name:  updateRequest.Name,
			Email: updateRequest.Email,
			Role:  updateRequest.Role,
		}
		if err := tx.Model(&user).Updates(user).Error; err != nil {
			return err
		}

		if err := tx.Where("user_id = ?", updateRequest.ID).Delete(&entity.Contacts{}).Error; err != nil {
			return err
		}

		for _, contact := range updateRequest.Contacts {
			contactEntity := entity.Contacts{
				UserID: user.ID,
				Phone:  contact.Phone,
				Bio:    contact.Bio,
			}
			if err := tx.Create(&contactEntity).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
