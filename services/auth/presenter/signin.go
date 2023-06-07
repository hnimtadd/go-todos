package presenter

type SignInResponse struct {
	Token string `json:"token"`
}

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
