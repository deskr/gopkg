package mailer

import (
	"gopkg.in/gomail.v2"
)

// Mailer interface
type Mailer interface {
	Send(email Email) (err error)
	AttachSentMailHandler(handler *SentMailHandler)
	RemoveSentMailHandler(handler *SentMailHandler)
}

// SentMailHandler listener for on sent email
type SentMailHandler func(email Email)

// HTML for HTML string
type HTML string

// Text for Text string
type Text string

// Body email body text and HTML
type Body struct {
	HTML HTML
	Text Text
}

type Address struct {
	Name    string
	Address string // email address
}

// format to RFC5322
// NB: Uses default gomail settings (charset, encoding)
func (a Address) toHeaderFormat() string {
	msg := gomail.NewMessage()
	return msg.FormatAddress(a.Address, a.Name)
}

// Email for building e-mails with text/template
type Email struct {
	From    Address
	To      Address
	ReplyTo *Address
	Subject string
	Body    Body
}
