package mailer

import (
	"gopkg.in/gomail.v2"
)

// gmailMailer for mailing through gmail relay service
type gmailMailer struct {
	fromEmail    string
	dialer       *gomail.Dialer
	sentHandlers []*SentMailHandler
}

// GMailConfig holds gmail configuration
type GMailConfig struct {
	// Gmail username
	Username string
	// Gmail password
	Password  string
	FromEmail string
}

// NewGMailer creates an new gmail mailer
func NewGMailer(config GMailConfig) Mailer {
	m := &gmailMailer{}

	m.dialer = gomail.NewPlainDialer("smtp-relay.gmail.com", 465,
		config.Username, config.Password)

	m.fromEmail = config.FromEmail

	return m
}

// AttachSentMailHandler attaches an handler to "OnSent" event
func (m *gmailMailer) AttachSentMailHandler(handler *SentMailHandler) {
	m.sentHandlers = append(m.sentHandlers, handler)
}

// RemoveSentMailHandler removes the handler
func (m *gmailMailer) RemoveSentMailHandler(handler *SentMailHandler) {
	for i, l := range m.sentHandlers {
		if l == handler {
			m.sentHandlers = append(m.sentHandlers[:i], m.sentHandlers[i+1:]...)
			break
		}
	}
}

// Send sends an email
func (m gmailMailer) Send(email Email) (err error) {

	fromEmail := email.From
	if fromEmail == "" {
		fromEmail = m.fromEmail
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", fromEmail)
	msg.SetHeader("To", email.To)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/plain", string(email.Body.Text))
	msg.AddAlternative("text/html", string(email.Body.HTML))

	err = m.dialer.DialAndSend(msg)
	if err != nil {
		return
	}

	for _, l := range m.sentHandlers {
		fn := *l
		fn(email)
	}

	return
}
