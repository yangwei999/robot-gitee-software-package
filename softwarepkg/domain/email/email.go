package email

type Email interface {
	Send(url, subject string) error
}
