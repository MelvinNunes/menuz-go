package sms

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gofiber/fiber/v2/log"

	"github.com/fiorix/go-smpp/smpp"
	"github.com/fiorix/go-smpp/smpp/pdu/pdufield"
	"github.com/fiorix/go-smpp/smpp/pdu/pdutext"
)

type SMSResult struct {
	PhoneNumber string
	Err         error
}

func SendSMS(phoneNumber string, message string) {
	send(phoneNumber, message)
}

func SendBulkSMS(phoneNumbers []string, message string) {
	for _, phoneNumber := range phoneNumbers {
		go send(phoneNumber, message)
	}
}

func send(phoneNumber string, message string) bool {
	host := os.Getenv("SMPP_HOST")
	port := os.Getenv("SMPP_PORT")
	user := os.Getenv("SMPP_USER")
	password := os.Getenv("SMPP_PASSWORD")

	if host == "" || port == "" {
		log.Panic("Please specify a host and port for SMPP_HOST and SMPP_PORT environment variables.")
	}
	if user == "" || password == "" {
		log.Panic("Please specify a user and password for SMPP_USER and SMPP_PASSWORD environment variables.")
	}

	tx := &smpp.Transmitter{
		Addr:   fmt.Sprintf("%v:%v", host, port),
		User:   user,
		Passwd: password,
	}

	// Create persistent connection, wait for the first status.
	conn := <-tx.Bind()
	if conn.Status() != smpp.Connected {
		log.Error("Failed to connect to SMPP server: ", conn.Error())
		return false
	}
	sm, err := tx.Submit(&smpp.ShortMessage{
		Src:           "VOLET",
		Dst:           fmt.Sprintf("258%v", phoneNumber),
		Text:          pdutext.Raw(replaceSpecialCharacters(message)),
		Register:      pdufield.NoDeliveryReceipt,
		SourceAddrTON: 5,
		SourceAddrNPI: 0,
	})
	if err != nil {
		log.Error(err)
		return false
	}
	log.Info("Message ID:", sm.RespID())
	tx.Close()
	return true
}

func replaceSpecialCharacters(s string) string {
	replacements := map[string]string{
		"é": "e",
		"è": "e",
		"à": "a",
		"á": "a",
		"ü": "u",
		"ö": "o",
		"ç": "c",
		"ã": "a",
		"õ": "o",
		"ó": "o",
		"ú": "u",
		"ù": "u",
		"ê": "e",
	}

	pattern := "["
	for accented := range replacements {
		pattern += accented
	}
	pattern += "]"

	re := regexp.MustCompile(pattern)
	replaced := re.ReplaceAllStringFunc(s, func(match string) string {
		return replacements[match]
	})
	return replaced
}
