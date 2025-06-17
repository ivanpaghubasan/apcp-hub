package structs

type CreateUserRequest struct {
	FirstName    string  `json:"firstName" binding:"required"`
	LastName     string  `json:"lastName" binding:"required"`
	DateOfBirth  string  `json:"dateOfBirth"`
	MobileNumber string  `json:"mobileNumber" binding:"required"`
	Gender       string  `json:"gender" binding:"required"`
	Position     *string `json:"position"`
	Username     string  `json:"username" binding:"required"`
	Password     string  `json:"password" binding:"required,min=8"`
}

type CreateUserResponse struct {
	UserID  string `json:"userId"`
	Message string `json:"message"`
}
