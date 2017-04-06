package mailer

import (
	"bytes"
	"os"
	"sync"
	"testing"
)

var mailgunTestConfig MailgunConfig

func init() {
	mailgunTestConfig.Domain = os.Getenv("MAILGUN_DOMAIN")
	mailgunTestConfig.PrivateAPIKey = os.Getenv("MAILGUN_PRIVATE_KEY")
	mailgunTestConfig.PublicAPIKey = os.Getenv("MAILGUN_PUBLIC_KEY")
}

func TestMailgunSendMail(t *testing.T) {
	if mailgunTestConfig.Domain == "" {
		t.SkipNow()
		return
	}

	m := NewMailgunMailer(mailgunTestConfig)

	err := m.Send(Email{
		To:      Address{Name: "Mike", Address: "mc@deskr.co"},
		From:    Address{Name: "Carina", Address: "carina@deskr.co"},
		ReplyTo: &Address{Name: "Martin", Address: "martin@deskr.co"},
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

func TestMailgunSendWithYAML(t *testing.T) {
	if mailgunTestConfig.Domain == "" {
		t.SkipNow()
		return
	}

	m := NewMailgunMailer(mailgunTestConfig)

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
		To:   Address{Name: "Mike", Address: "mc@deskr.co"},
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
