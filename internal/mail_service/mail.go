package mail_service

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/pkg/errors"
)

type Mail struct {
	servername string
	username   string
	password   string
}

func (m *Mail) createClient() (*smtp.Client, error) {
	host, _, _ := net.SplitHostPort(m.servername)
	auth := smtp.PlainAuth("", m.username, m.password, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Contract dial key
	conn, err := tls.Dial("tcp", m.servername, tlsconfig)
	if err != nil {
		return nil, errors.Wrap(err, "[tls.Dial]")
	}

	// Create Client
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return nil, errors.Wrap(err, "[smtp.NewClient]")
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		return nil, errors.Wrap(err, "[c.Auth]")
	}

	// From
	if err = c.Mail(m.username); err != nil {
		return nil, errors.Wrapf(err, "[c.Mail(%v)]", m.username)
	}

	return c, nil
}

func NewMail(servername string, username string, password string) (*Mail, error) {

	return &Mail{
		servername: servername,
		username:   username,
		password:   password,
	}, nil
}

type SendMailIn struct {
	To      string
	Subject string
	Body    string
}

func (m *Mail) SendMail(params *SendMailIn) error {
	from := mail.Address{"", m.username}
	to := mail.Address{"", params.To}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = params.Subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + params.Body

	c, err := m.createClient()
	if err != nil {
		return errors.Wrap(err, "[m.createClient]")
	}

	if err := c.Rcpt(to.Address); err != nil {
		return errors.Wrapf(err, "[m.c.Rcpt(%v)]", to.Address)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		return errors.Wrap(err, "[m.c.Data]")
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return errors.Wrap(err, "[w.Write]")
	}

	err = w.Close()
	if err != nil {
		return errors.Wrap(err, "[w.Close]")
	}

	err = c.Quit()
	if err != nil {
		return errors.Wrap(err, "[c.Quit]")
	}

	return nil
}
