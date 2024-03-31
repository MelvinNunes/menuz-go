package twilio

import (
	"log"
	"os"

	"github.com/twilio/twilio-go"
)

func twilioClient() *twilio.RestClient {
	accountSid := os.Getenv("TWILIO_SID")
	if accountSid == "" {
		log.Panic("TWILIO_SID must be set in environment variables")
	}

	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	if authToken == "" {
		log.Panic("TWILIO_AUTH_TOKEN must be set in environment variables")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return client
}
