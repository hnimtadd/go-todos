package presenter

type SignUpRepsonse struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Limit    int    `json:"limit"`
}

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Limit    int    `json:"limit"`
}
