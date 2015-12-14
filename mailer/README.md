# mailer

Implemented providers:

- Fake (writes mail output to stdout)
- Gmail
- Mailgun

Example:
```go
m := NewFakeMailer()

err := m.Send(Email{
    To:      "mike@deskr.co",
    From:    "carina@deskr.co",
    Subject: "Something important",
    Body: Body{
        Text: "Hello you",
        HTML: "<strong>Hello you</strong>",
    },
})
if err != nil {
    panic(err)
}

```