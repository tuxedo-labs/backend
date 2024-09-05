package request

type Contact struct {
	Phone *string `json:"phone"`
	Bio   *string `json:"bio"`
}

type UserProfile struct {
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Role      string  `json:"role"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	Contacts  Contact `json:"contacts"`
}

type UpdateUserProfileRequest struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Contacts Contact `json:"contacts"`
}
