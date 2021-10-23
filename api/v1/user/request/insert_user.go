package request

import "go-hexagonal/business/user"

//InsertUserRequest create User request payload
type InsertUserRequest struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
}

//ToUpsertUserSpec convert into User.UpsertUserSpec object
func (req *InsertUserRequest) ToUpsertUserSpec() *user.InsertUserSpec {

	var insertUserSpec user.InsertUserSpec

	insertUserSpec.Firstname = req.Firstname
	insertUserSpec.Lastname = req.Lastname
	insertUserSpec.Email = req.Email
	insertUserSpec.Username = req.Username
	insertUserSpec.Password = req.Password
	insertUserSpec.Phone = req.Phone

	return &insertUserSpec
}
