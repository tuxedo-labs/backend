package entity

type Contacts struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"index"`
	Phone  *string `json:"phone"`
	Bio    *string `json:"bio"`

	User Users `gorm:"foreignKey:UserID"`
}