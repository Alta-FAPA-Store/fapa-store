package response

import (
	"go-hexagonal/api/paginator"
	"go-hexagonal/business/user"
)

type getAllUserResponse struct {
	Meta  paginator.Meta    `json:"meta"`
	Users []GetUserResponse `json:"users"`
}

//NewGetAllUserResponse construct GetAllUserResponse
func NewGetAllUserResponse(users []user.User, page int, rowPerPage int) getAllUserResponse {

	var (
		lenUsers = len(users)
	)

	getAllUserResponse := getAllUserResponse{}
	getAllUserResponse.Meta.BuildMeta(lenUsers, page, rowPerPage)

	for index, value := range users {
		if index == getAllUserResponse.Meta.RowPerPage {
			continue
		}

		var getUserResponse GetUserResponse

		getUserResponse.ID = value.ID
		getUserResponse.Firstname = value.Firstname
		getUserResponse.Lastname = value.Lastname
		getUserResponse.Email = value.Email
		getUserResponse.Phone = value.Phone
		getUserResponse.Username = value.Username

		getAllUserResponse.Users = append(getAllUserResponse.Users, getUserResponse)
	}

	if getAllUserResponse.Users == nil {
		getAllUserResponse.Users = []GetUserResponse{}
	}

	return getAllUserResponse
}
