package services

import (
	"errors"
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
	var contact request.Contacts

	if user.Contacts != (entity.Contacts{}) {
		contact = request.Contacts{
			Phone: &user.Contacts.Phone,
			Bio:   &user.Contacts.Bio,
		}
	}

	profile := request.UserProfile{
		Name:      user.Name,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		Verify:    user.Verify,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02"),
		Contacts:  contact,
	}

	return profile, nil
}

func UpdateUserProfile(updateRequest request.UpdateUserProfileRequest) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// Update user details
		user := entity.Users{
			ID:    updateRequest.ID,
			Name:  updateRequest.Name,
			Email: updateRequest.Email,
		}
		if err := tx.Model(&user).Updates(user).Error; err != nil {
			return err
		}

		// Fetch existing contact
		var existingContact entity.Contacts
		if err := tx.Where("user_id = ?", updateRequest.ID).First(&existingContact).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		// Prepare contact for update
		contact := entity.Contacts{
			UserID: updateRequest.ID,
		}

		// Only assign values if they are not nil
		if updateRequest.Contacts.Phone != nil {
			contact.Phone = *updateRequest.Contacts.Phone // Dereference safely
		}
		if updateRequest.Contacts.Bio != nil {
			contact.Bio = *updateRequest.Contacts.Bio // Dereference safely
		}

		// Create or update the contact
		if existingContact.ID == 0 {
			if err := tx.Create(&contact).Error; err != nil {
				return err
			}
		} else {
			if err := tx.Model(&existingContact).Updates(contact).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
