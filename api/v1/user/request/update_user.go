package request

import "go-hexagonal/business/user"

//UpdateUserRequest update User request payload
type UpdateUserRequest struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Version   int    `json:"version"`
}

func (req *UpdateUserRequest) ToUpsertUserSpec() *user.UpdateUserRequest {

	var updateUserRequest user.UpdateUserRequest

	updateUserRequest.Firstname = req.Firstname
	updateUserRequest.Lastname = req.Lastname
	updateUserRequest.Email = req.Email
	updateUserRequest.Username = req.Username
	updateUserRequest.Phone = req.Phone
	updateUserRequest.Version = req.Version

	return &updateUserRequest
}
