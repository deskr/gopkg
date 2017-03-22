package mailer

import (
	"fmt"

	"github.com/mailgun/mailgun-go"
)

// mailgunMailer for mailing through mailgun service
type mailgunMailer struct {
	mg           mailgun.Mailgun
	sentHandlers []*SentMailHandler
}

// MailgunConfig hold domain, apiKey and publicApiKey
type MailgunConfig struct {
	Domain        string
	PrivateAPIKey string
	PublicAPIKey  string
}

// NewMailgunMailer creates an new mailgun mailer
func NewMailgunMailer(config MailgunConfig) Mailer {
	m := &mailgunMailer{}

	m.mg = mailgun.NewMailgun(config.Domain,
		config.PrivateAPIKey, config.PublicAPIKey)

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
	msg := m.mg.NewMessage(
		"",
		email.Subject,
		string(email.Body.Text),
		"",
	)
	msg.AddHeader("From", fmt.Sprintf("%s <%s>", email.From.Name, email.From.Address))
	msg.AddHeader("To", fmt.Sprintf("%s <%s>", email.To.Name, email.To.Address))
	if email.ReplyTo != nil {
		msg.AddHeader("Reply-To", fmt.Sprintf("%s <%s>", email.ReplyTo.Name, email.ReplyTo.Address))
	}
	msg.SetHtml(string(email.Body.HTML))

	_, _, err = m.mg.Send(msg)
	if err != nil {
		return
	}

	for _, l := range m.sentHandlers {
		fn := *l
		fn(email)
	}

	return
}
