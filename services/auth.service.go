package services

import (
	"crypto/rand"
	"fmt"
	"time"
	"tuxedo/database"
	"tuxedo/lib"
	"tuxedo/middleware"
	"tuxedo/models/entity"
	"tuxedo/models/request"
	"tuxedo/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
)

func ValidateLogin(loginRequest *request.LoginRequest) error {
	validate := validator.New()
	return validate.Struct(loginRequest)
}

func GetUserByEmail(email string) (*entity.Users, error) {
	var user entity.Users
	err := database.DB.First(&user, "email = ?", email).Error
	return &user, err
}

func GenerateJWTToken(user *entity.Users) (string, error) {
	claims := jwt.MapClaims{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
		"role":  "member",
	}

	if user.Role == "admin" {
		claims["role"] = "admin"
	}

	return utils.GenerateToken(&claims)
}

func ValidateRegister(registerRequest *request.RegisterRequest) error {
	validate := validator.New()
	return validate.Struct(registerRequest)
}

func HashAndStoreUser(registerRequest *request.RegisterRequest) (string, error) {
	var existingUser entity.Users
	if err := database.DB.First(&existingUser, "email = ?", registerRequest.Email).Error; err == nil {
		return "", fmt.Errorf("user with email %s already exists", registerRequest.Email)
	}

	hashedPassword, err := middleware.HashPassword(registerRequest.Password)
	if err != nil {
		return "", err
	}

	newUser := entity.Users{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
		Role:     "member",
		Verify:   false,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return "", err
	}

	token, err := generateVerificationToken()
	if err != nil {
		return "", err
	}

	verifyToken := entity.VerifyToken{
		Token:  token,
		UserID: newUser.ID,
	}
	if err := database.DB.Create(&verifyToken).Error; err != nil {
		return "", err
	}

	err = lib.SendVerificationEmail(newUser.Email, token)
	if err != nil {
		return "", err
	}
	// change token to status if in prod
	return token, nil
}

func UpdateUser(user *entity.Users) error {
	return database.DB.Save(user).Error
}

func DeleteVerifyToken(tokenID uint) error {
	return database.DB.Delete(&entity.VerifyToken{}, tokenID).Error
}

func GetVerifyToken(token string) (*entity.VerifyToken, error) {
	var verifyToken entity.VerifyToken
	if err := database.DB.Where("token = ?", token).First(&verifyToken).Error; err != nil {
		return nil, err
	}
	return &verifyToken, nil
}

func generateVerificationToken() (string, error) {
	token := make([]byte, 4)
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	num := uint32(token[0])<<24 | uint32(token[1])<<16 | uint32(token[2])<<8 | uint32(token[3])

	return fmt.Sprintf("%06d", num%1000000), nil
}

func DeleteVerifyTokenByUserID(userID uint) error {
	return database.DB.Where("user_id = ?", userID).Delete(&entity.VerifyToken{}).Error
}

func GenerateAndSendVerificationToken(user *entity.Users) error {
	token, err := generateVerificationToken()
	if err != nil {
		return err
	}

	verifyToken := entity.VerifyToken{
		Token:  token,
		UserID: user.ID,
	}

	if err := database.DB.Create(&verifyToken).Error; err != nil {
		return err
	}

	err = lib.SendVerificationEmail(user.Email, token)
	if err != nil {
		return err
	}

	return nil
}

func AuthenticateUser(email, password string) (*entity.Users, error) {
	var user entity.Users
	err := database.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return nil, err
	}

	if !middleware.CheckPassword(user.Password, password) {
		return nil, fmt.Errorf("invalid password")
	}

	return &user, nil
}
