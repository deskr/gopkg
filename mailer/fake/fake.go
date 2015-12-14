package fake

import (
	"fmt"

	"github.com/deskr/gopkg/mailer"
)

// Mailer for mailing through gmail relay service
type Mailer struct {
	sentHandlers []*mailer.SentMailHandler
}

// NewMailer creates an new gmail mailer
func NewMailer() mailer.Mailer {
	return &Mailer{}
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

	fmt.Printf("Fake email sent: %+v", email)

	for _, l := range m.sentHandlers {
		fn := *l
		fn(email)
	}

	return
}
