package notification

import "fmt"

type EmailNotification struct{}

func (e *EmailNotification) Send(to string, message string) string {
	return fmt.Sprintf("Email enviado a %s con mensaje: %s", to, message)
}
