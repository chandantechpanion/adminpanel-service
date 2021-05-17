package login

type Repository interface {
	SignIn(*Login) (string, error)
}
