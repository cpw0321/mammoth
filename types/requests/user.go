package requests

// UserRequest ...
type UserRequest struct {
	UserName string `json:"user_name"`
	//Telephone string `json:"telephone"`
	Password string `json:"password"`
}
