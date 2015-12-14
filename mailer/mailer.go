package mailer

// Mailer interface
type Mailer interface {
	Send(email Email) (err error)
	AttachSentMailHandler(listener *SentMailHandler)
	RemoveSentMailHandler(listener *SentMailHandler)
}

// SentMailHandler for on sent email event
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

// Email for building e-mails with text/template
type Email struct {
	From    string
	To      string
	Subject string
	Body    Body
}
