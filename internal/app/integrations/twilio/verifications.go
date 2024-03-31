package twilio

import (
	"fmt"
	"log"
	"os"

	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

const (
	VERIFICATION_APPROVED = "approved"
	VERIFICATION_
)

func SendPhoneVerificationCode(phoneCode string, phoneNumber string) bool {
	client := twilioClient()
	verificationSID := os.Getenv("TWILIO_VERIFY_SID")
	if verificationSID == "" {
		log.Println("TWILIO_VERIFY_SID must be set in environment variables")
	}

	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneCode + phoneNumber)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(verificationSID, params)
	if err != nil {
		log.Println(err.Error())
	} else {
		if resp.Sid != nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func SendMailVerificationCode(email string) bool {
	client := twilioClient()
	verificationSID := os.Getenv("TWILIO_VERIFY_SID")
	if verificationSID == "" {
		log.Println("TWILIO_VERIFY_SID must be set in environment variables")
	}
	params := &verify.CreateVerificationParams{}
	params.SetTo(email)
	params.SetChannel("email")

	resp, err := client.VerifyV2.CreateVerification(verificationSID, params)
	if err != nil {
		log.Println(err.Error())
	} else {
		if resp.Sid != nil {
			return true
		} else {
			return false
		}
	}
	return false
}

func IsVerificationCodeValid(code string, receiver string) bool {
	client := twilioClient()
	verificationSID := os.Getenv("TWILIO_VERIFY_SID")
	if verificationSID == "" {
		log.Panic("TWILIO_VERIFY_SID must be set in environment variables")
	}

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(receiver)
	params.SetCode(code)

	resp, err := client.VerifyV2.CreateVerificationCheck(verificationSID, params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Status != nil {
			return *resp.Status == VERIFICATION_APPROVED
		} else {
			return false
		}
	}
	return false
}
