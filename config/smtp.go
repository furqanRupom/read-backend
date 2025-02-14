package config

import (
	"net/smtp"
	"net/url"
	"time"

	gomail "gopkg.in/mail.v2"
)

type SMTPConfig struct {
	Dialer   *gomail.Dialer
	Username string
	BaseURL  url.URL
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}

func createSMTPConfig() SMTPConfig {
	host := getEnvWithoutParser("SMTP_HOST", false)
	port := getEnv("SMTP_PORT", false, parseInt)
	user := getEnvWithoutParser("SMTP_USER", false)
	password := getEnvWithoutParser("SMTP_PASSWORD", false)
	useSSL := getEnv("SMTP_USE_SSL", false, parseBool)
	baseURL := getEnv("EMAILS_CONFIRM_BASE_URL", false, url.Parse)
	dialer := &gomail.Dialer{
		Host: host,
		Port: port,
		Auth: unencryptedAuth{
			smtp.PlainAuth(
				"",
				user,
				password,
				host,
			),
		},
		SSL:          useSSL,
		Timeout:      10 * time.Second,
		RetryFailure: true,
	}
	return SMTPConfig{
		Dialer:   dialer,
		Username: user,
		BaseURL:  baseURL,
	}
}
