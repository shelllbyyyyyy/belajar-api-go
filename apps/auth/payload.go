package auth

type updateUserPayload struct {
	Id       string
	Username *string
	Email    *string
	Password *string
}