package auth

func toUserResponse(model *User) userResponse {
	return userResponse{
		Id:        model.Id,
		Username:  model.Username,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}