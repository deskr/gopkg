package mailgun

import (
	"github.com/deskr/gopkg/mailer"

	"github.com/mailgun/mailgun-go"
)

// Mailer for mailing through mailgun service
type Mailer struct {
	mg           mailgun.Mailgun
	from         string
	sentHandlers []*mailer.SentMailHandler
}

// Config hold domain, apiKey and publicApiKey
type Config struct {
	Domain        string
	PrivateAPIKey string
	PublicAPIKey  string
}

// NewMailer creates an new mailgun mailer
func NewMailer(config Config, options ...func(*Mailer)) mailer.Mailer {
	m := &Mailer{}

	m.mg = mailgun.NewMailgun(config.Domain,
		config.PrivateAPIKey, config.PublicAPIKey)

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

	mail := m.mg.NewMessage(
		from,
		email.Subject,
		string(email.Body.Text),
		email.To,
	)
	mail.SetHtml(string(email.Body.HTML))

	_, _, err = m.mg.Send(mail)
	if err != nil {
		return
	}

	for _, l := range m.sentHandlers {
		fn := *l
		fn(email)
	}

	return
}
