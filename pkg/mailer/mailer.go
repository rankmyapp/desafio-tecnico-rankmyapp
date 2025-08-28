package mailer

import (
    "fmt"
    "net/smtp"
    "strings"

    "desafio-tecnico/internal/config"
)

type Mailer struct {
    host string
    port int
    user string
    pass string
    from string
}

func New(cfg *config.Config) *Mailer {
    return &Mailer{
        host: cfg.SMTPHost,
        port: cfg.SMTPPort,
        user: cfg.SMTPUser,
        pass: cfg.SMTPPass,
        from: cfg.SMTPFrom,
    }
}

func (m *Mailer) Send(to, subject, body string) error {
    addr := fmt.Sprintf("%s:%d", m.host, m.port)
    msg := buildMessage(m.from, to, subject, body)

    var auth smtp.Auth
    if strings.TrimSpace(m.user) != "" {
        auth = smtp.PlainAuth("", m.user, m.pass, m.host)
    } else {
        auth = nil
    }
    return smtp.SendMail(addr, auth, m.from, []string{to}, []byte(msg))
}

func buildMessage(from, to, subject, body string) string {
    headers := []string{
        "From: " + from,
        "To: " + to,
        "Subject: " + subject,
        "MIME-Version: 1.0",
        "Content-Type: text/plain; charset=UTF-8"   }
    return strings.Join(headers, "\r\n") + "\r\n\r\n" + body
}