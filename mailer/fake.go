package mailer

import (
	"fmt"
)

type fakeMailer struct {
	sentHandlers []*SentMailHandler
}

// NewFakeMailer creates a new fake email sender
func NewFakeMailer() Mailer {
	return &fakeMailer{}
}

// AttachSentMailHandler attaches an handler to "OnSent" event
func (m *fakeMailer) AttachSentMailHandler(handler *SentMailHandler) {
	m.sentHandlers = append(m.sentHandlers, handler)
}

// RemoveSentMailHandler removes the handler
func (m *fakeMailer) RemoveSentMailHandler(handler *SentMailHandler) {
	for i, l := range m.sentHandlers {
		if l == handler {
			m.sentHandlers = append(m.sentHandlers[:i], m.sentHandlers[i+1:]...)
			break
		}
	}
}

// Send sends an email
func (m fakeMailer) Send(email Email) (err error) {

	fmt.Printf("Fake email sent: %+v", email)

	for _, l := range m.sentHandlers {
		fn := *l
		fn(email)
	}

	return
}
