package entity

type Contacts struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	Phone  string `json:"phone"`
	Bio    string `json:"bio"`
}
