# mailer

Implemented providers:

- Fake (writes mail output to stdout)
- Gmail
- Mailgun

Example:
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