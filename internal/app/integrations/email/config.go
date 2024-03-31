package mailer

import (
	"log"

	"github.com/go-mail/mail"
)

func SendMail() {
	// Set up authentication information
	from := "payment@voletapp.com"
	password := "FjUD.suy$)B^C9u"
	host := "smtp-mail.outlook.com"
	port := 587

	// Set up message
	subject := "Testing email sender?"
	body := "Testing"

	m := mail.NewMessage()

	m.SetHeader("From", from)

	m.SetHeader("To", "melvinfulana@gmail.com")

	// m.SetAddressHeader("Cc", "oliver.doe@example.com", "Oliver")

	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// m.Attach("lolcat.jpg")

	d := mail.NewDialer(host, port, from, password)
	if err := d.DialAndSend(m); err != nil {
		log.Panic(err)
	}
}
