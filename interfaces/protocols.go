package interfaces


type Gmailer interface {
	Send(from, to, subject, msg string) error
}
