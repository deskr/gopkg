package mailer

import (
	"bytes"
	ht "html/template"
	"io/ioutil"
	tt "text/template"

	"gopkg.in/yaml.v2"
)

// EmailTemplate represents a simple email with html and text body
// in Go template language HTML and Text
type EmailTemplate struct {
	Subject string
	Body    struct {
		Text tt.Template
		HTML ht.Template
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
func (e EmailTemplate) Execute(textData interface{}, htmlData interface{}) (b Body, err error) {
	var buf bytes.Buffer

	if err = e.Body.Text.Execute(&buf, textData); err != nil {
		return
	}

	b.Text = Text(buf.String())
	buf.Reset()

	if err = e.Body.HTML.Execute(&buf, htmlData); err != nil {
		return
	}

	b.HTML = HTML(buf.String())
	return
}

// ParseYAML gets an email template from file
func ParseYAML(file string) (t EmailTemplate, err error) {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	var yt emailTemplateYAML
	if err = yaml.Unmarshal([]byte(data), &yt); err != nil {
		return
	}

	t = EmailTemplate{
		Subject: yt.Subject,
	}

	btt, err := tt.New("").Parse(yt.Body.Text)
	if err != nil {
		return
	}
	t.Body.Text = *btt

	bht, err := ht.New("").Parse(yt.Body.HTML)
	if err != nil {
		return
	}
	t.Body.HTML = *bht

	return
}
