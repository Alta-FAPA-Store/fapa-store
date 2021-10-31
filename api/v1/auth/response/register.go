package response

//Login response payload
type RegisterResponse struct {
	ValidationId string `json:"validation_id"`
}

//NewLoginResponse construct LoginResponse
func NewRegisterResponse(id string) *RegisterResponse {
	var RegisterResponse RegisterResponse

	RegisterResponse.ValidationId = id

	return &RegisterResponse
}
