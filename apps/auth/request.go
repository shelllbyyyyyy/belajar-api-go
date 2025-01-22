package auth

type registerUserSchema struct {
	Username string `json:"username" validate:"required, min=4"`
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}

type loginUserSchema struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8"`
}

type updateUserSchema struct {
	Username *string `json:"username" validate:"optional, min=4"`
	Email    *string `json:"email" validate:"optional, email"`
	Password *string `json:"password" validate:"optional, min=8"`
}

type tokenSchema struct {
	Id string
}