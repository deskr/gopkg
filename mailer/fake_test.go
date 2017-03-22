package mailer

import (
	"bytes"
	"sync"
	"testing"
)

func TestSendMail(t *testing.T) {
	m := NewFakeMailer()

	err := m.Send(Email{
		To:      Address{Name: "Mike", Address: "mike@deskr.co"},
		From:    Address{Name: "Carina", Address: "carina@deskr.co"},
		Subject: "Something important",
		Body: Body{
			Text: "Hello you",
			HTML: "<strong>Hello you</strong>",
		},
	})

	if err != nil {
		t.Errorf("Failed to send mail: %s", err)
	}
}

func TestAddRemoveSentHandler(t *testing.T) {
	m := &fakeMailer{}

	wg := sync.WaitGroup{}

	wg.Add(1)
	h := SentMailHandler(func(email Email) {
		wg.Done()
	})

	m.AttachSentMailHandler(&h)
	if len(m.sentHandlers) != 1 {
		t.Errorf("Expected 1 listener, got: %d", len(m.sentHandlers))
		return
	}

	m.Send(Email{
		To:      Address{Name: "Mike", Address: "mike@deskr.co"},
		From:    Address{Name: "Carina", Address: "carina@deskr.co"},
		Subject: "Something important",
		Body: Body{
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

func TestSendWithYAML(t *testing.T) {
	m := NewFakeMailer()

	tmpl, err := ParseYAML(bytes.NewReader([]byte(
		`subject: Something important
body:
    text: Hello {{.Who}}
    html: <strong>Hello {{.Who}}</strong>
`)))

	if err != nil {
		t.Errorf("Failed to parse yaml: %s", err)
		return
	}

	email := Email{
		To:   Address{Name: "Mike", Address: "mike@deskr.co"},
		From: Address{Name: "Carina", Address: "carina@deskr.co"},
	}

	data := struct {
		Who string
	}{
		Who: "Someone",
	}

	td := EmailTemplateData{}
	td.Body.HTML = data
	td.Body.Text = data

	err = tmpl.Execute(&email, td)
	if err != nil {
		t.Errorf("Failed to execute template data: %s", err)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	h := SentMailHandler(func(email Email) {
		if email.Body.Text != "Hello Someone" {
			t.Errorf("Unexpected Body.Text for sent email: %s", email.Body.Text)
		}
		wg.Done()
	})
	m.AttachSentMailHandler(&h)

	err = m.Send(email)

	if err != nil {
		t.Errorf("Failed to send mail: %s", err)
	}

	wg.Wait()
}
