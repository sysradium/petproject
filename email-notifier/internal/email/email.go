package email

type Message struct {
	To      string
	Subject string
	From    string
	Body    string
}

type Sender interface {
	Send(m Message) error
}
