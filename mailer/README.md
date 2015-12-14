# mailer

Implemented providers:

- Fake (writes mail output to stdout)
- Gmail
- Mailgun

### Examples

Sending an email:
```go
m := mailer.NewFakeMailer()

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
    panic(err)
}

```


Sending an email with YAML template:
```go
m := mailer.NewFakeMailer()

tmpl, err := mailer.ParseYAML([]byte(
    `subject: Something important
body:
    text: Hello {{.Who}}
    html: <strong>Hello {{.Who}}</strong>
`))
if err != nil {
    panic(err)
}

data := struct {
    Who string
}{
    Who: "You",
}

body, err := tmpl.Execute(data, data)
if err != nil {
    panic(err)
}

err = m.Send(mailer.Email{
    To:      "mike@deskr.co",
    From:    "carina@deskr.co",
    Subject: "Something important",
    Body:    body,
})
if err != nil {
    panic(err)
}
```


Sending an email with "on sent" email handler:
```go
m := mailer.NewFakeMailer()

wg := sync.WaitGroup{}
wg.Add(1)

h := mailer.SentMailHandler(func(email Email) {
    wg.Done()
})
m.AttachSentMailHandler(&h)

go func() {
    m.Send(mailer.Email{
        To:      "mike@deskr.co",
        From:    "carina@deskr.co",
        Subject: "Something important",
        Body: Body{
            Text: "Hello you",
            HTML: "<strong>Hello you</strong>",
        },
    })
}()

wg.Wait()
```