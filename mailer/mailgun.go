package mailer

import (
	"github.com/mailgun/mailgun-go"
)

// mailgunMailer for mailing through mailgun service
type mailgunMailer struct {
	mg           mailgun.Mailgun
	fromEmail    string
	sentHandlers []*SentMailHandler
}

// MailgunConfig hold domain, apiKey and publicApiKey
type MailgunConfig struct {
	Domain        string
	FromEmail     string
	PrivateAPIKey string
	PublicAPIKey  string
}

// NewMailgunMailer creates an new mailgun mailer
func NewMailgunMailer(config MailgunConfig) Mailer {
	m := &mailgunMailer{}

	m.mg = mailgun.NewMailgun(config.Domain,
		config.PrivateAPIKey, config.PublicAPIKey)

	m.fromEmail = config.FromEmail

	return m
}

// AttachSentMailHandler attaches an handler to "OnSent" event
func (m *mailgunMailer) AttachSentMailHandler(handler *SentMailHandler) {
	m.sentHandlers = append(m.sentHandlers, handler)
}

// RemoveSentMailHandler removes the handler
func (m *mailgunMailer) RemoveSentMailHandler(handler *SentMailHandler) {
	for i, l := range m.sentHandlers {
		if l == handler {
			m.sentHandlers = append(m.sentHandlers[:i], m.sentHandlers[i+1:]...)
			break
		}
	}
}

// Send sends an email
func (m mailgunMailer) Send(email Email) (err error) {

	fromEmail := email.From
	if fromEmail == "" {
		fromEmail = m.fromEmail
	}

	mail := m.mg.NewMessage(
		fromEmail,
		email.Subject,
		string(email.Body.Text),
		email.To,
	)
	if email.ReplyTo != "" {
		mail.AddHeader("Reply-To", email.ReplyTo)
	}
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
