package gmail

import (
	"github.com/deskr/gopkg/mailer"

	"gopkg.in/gomail.v2"
)

// Mailer for mailing through gmail relay service
type Mailer struct {
	from         string
	dialer       *gomail.Dialer
	sentHandlers []*mailer.SentMailHandler
}

// Credentials holds the gmail username & password
type Credentials struct {
	Username string
	Password string
}

// NewMailer creates an new gmail mailer
func NewMailer(cred Credentials, options ...func(*Mailer)) mailer.Mailer {
	m := &Mailer{}

	m.dialer = gomail.NewPlainDialer("smtp-relay.gmail.com", 465,
		cred.Username, cred.Password)

	// apply settings/options
	for _, o := range options {
		o(m)
	}

	return m
}

// FromEmail set the address from field
func FromEmail(from string) func(*Mailer) {
	return func(m *Mailer) {
		m.from = from
		return
	}
}

// AttachSentMailHandler attaches an handler to "OnSent" event
func (m *Mailer) AttachSentMailHandler(handler *mailer.SentMailHandler) {
	m.sentHandlers = append(m.sentHandlers, handler)
}

// RemoveSentMailHandler removes the handler
func (m *Mailer) RemoveSentMailHandler(handler *mailer.SentMailHandler) {
	for i, l := range m.sentHandlers {
		if l == handler {
			m.sentHandlers = append(m.sentHandlers[:i], m.sentHandlers[i+1:]...)
			break
		}
	}
}

// Send sends an email
func (m Mailer) Send(email mailer.Email) (err error) {

	from := email.From
	if from == "" {
		from = m.from
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", from)
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
