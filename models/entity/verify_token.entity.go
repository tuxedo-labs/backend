package entity

type VerifyToken struct {
	ID     uint   `gorm:"primaryKey"`
	Token  string `json:"token"`
	UserID uint   `gorm:"index"`
	User   Users  `gorm:"foreignKey:UserID"`
}
