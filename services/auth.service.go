package services

import (
	"fmt"
	"time"
	"tuxedo/database"
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
	err := database.DB.Debug().First(&user, "email = ?", email).Error
	return &user, err
}

func GenerateJWTToken(user *entity.Users) (string, error) {
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Minute * 2).Unix(),
		"role":  "user",
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

func HashAndStoreUser(registerRequest *request.RegisterRequest) error {
	var existingUser entity.Users
	err := database.DB.Debug().First(&existingUser, "email = ?", registerRequest.Email).Error
	if err == nil {
		return fmt.Errorf("user with email %s already exists", registerRequest.Email)
	}

	hashedPassword, err := middleware.HashPassword(registerRequest.Password)
	if err != nil {
		return err
	}

	newUser := entity.Users{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
		Role:     "member",
	}

	return database.DB.Create(&newUser).Error
}
