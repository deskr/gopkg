package fake

import (
	"sync"
	"testing"

	"github.com/deskr/gopkg/mailer"
)

type SentMailListener func(email mailer.Email)

func (l SentMailListener) OnSent(email mailer.Email) {
	l(email)
}

func TestSendMail(t *testing.T) {
	m := NewMailer()

	err := m.Send(mailer.Email{
		To:      "mike@deskr.co",
		From:    "carina@deskr.co",
		Subject: "Something important",
		Body: mailer.Body{
			Text: "Hello you",
			HTML: "<strong>Hello you</strong>",
		},
	})

	if err != nil {
		t.Errorf("Failed to send mail: %s", err)
	}
}

func TestAddRemoveSentHandler(t *testing.T) {
	m := &Mailer{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	h := mailer.SentMailHandler(func(email mailer.Email) {
		wg.Done()
	})

	m.AttachSentMailHandler(&h)
	if len(m.sentHandlers) != 1 {
		t.Errorf("Expected 1 listener, got: %d", len(m.sentHandlers))
		return
	}

	m.Send(mailer.Email{
		To:      "mike@deskr.co",
		From:    "carina@deskr.co",
		Subject: "Something important",
		Body: mailer.Body{
			Text: "Hello you",
			HTML: "<strong>Hello you</strong>",
		},
	})

	m.RemoveSentMailHandler(&h)
	if len(m.sentHandlers) != 0 {
		t.Errorf("Expected 0 listeners, got: %d", len(m.sentHandlers))
		return
	}

	wg.Wait()
}
