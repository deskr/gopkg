package mailer

import (
	"bytes"
	ht "html/template"
	"io"
	"io/ioutil"
	tt "text/template"

	"gopkg.in/yaml.v2"
)

// EmailTemplate represents a simple email with html and text body
type EmailTemplate struct {
	Subject *tt.Template
	Body    struct {
		Text *tt.Template
		HTML *ht.Template
	}
}

// EmailData is the data to execute with the template
type EmailData struct {
	Subject interface{}
	Body    struct {
		Text interface{}
		HTML interface{}
	}
}

type emailTemplateYAML struct {
	Subject string `yaml:"subject"`
	Body    struct {
		Text string `yaml:"text"`
		HTML string `yaml:"html"`
	}
}

// Execute the body templates
func (e EmailTemplate) Execute(d EmailData) (email Email, err error) {
	var buf bytes.Buffer

	// Execute subject
	if err = e.Subject.Execute(&buf, d.Subject); err != nil {
		return
	}
	email.Subject = buf.String()
	buf.Reset()

	// Execute body text
	if err = e.Body.Text.Execute(&buf, d.Body.Text); err != nil {
		return
	}
	email.Body.Text = Text(buf.String())
	buf.Reset()

	// Execute body HTML
	if err = e.Body.HTML.Execute(&buf, d.Body.HTML); err != nil {
		return
	}
	email.Body.HTML = HTML(buf.String())

	return
}

// ParseYAML reads and parses an email template
func ParseYAML(r io.Reader) (t EmailTemplate, err error) {

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}

	var yt emailTemplateYAML
	if err = yaml.Unmarshal(b, &yt); err != nil {
		return
	}

	t = EmailTemplate{}

	// Parse subject
	t.Subject, err = tt.New("").Parse(yt.Subject)
	if err != nil {
		return
	}

	// Parse body text
	t.Body.Text, err = tt.New("").Parse(yt.Body.Text)
	if err != nil {
		return
	}

	// Parse body HTML
	t.Body.HTML, err = ht.New("").Parse(yt.Body.HTML)
	if err != nil {
		return
	}

	return
}
