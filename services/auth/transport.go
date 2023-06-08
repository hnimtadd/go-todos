package auth

type AuthTransport interface {
	SignIn() error
	SignUp() error
}
