package useradapter

type UserAdapter interface {
	GetEmail(string) (string, error)
}
