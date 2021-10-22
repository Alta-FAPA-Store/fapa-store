package response

import (
	"go-hexagonal/business/user"
)

//GetUserResponse Get user by ID response payload
type GetUserResponse struct {
	ID        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
}

//NewGetUserResponse construct GetUserResponse
func NewGetUserResponse(user user.User) *GetUserResponse {
	var getUserResponse GetUserResponse

	getUserResponse.ID = user.ID
	getUserResponse.Firstname = user.Firstname
	getUserResponse.Lastname = user.Lastname
	getUserResponse.Email = user.Email
	getUserResponse.Username = user.Username

	return &getUserResponse
}
