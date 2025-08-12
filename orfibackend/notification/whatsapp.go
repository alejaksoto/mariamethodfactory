package notification

import "fmt"

type WhatsAppNotification struct{}

func (w *WhatsAppNotification) Send(recipient string, message string) string {
	return fmt.Sprintf("WhatsApp enviado a %s con mensaje: %s", recipient, message)
}
