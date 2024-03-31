package mpesa

import (
	"fmt"
	"log"
)

const (
	SuccessfulRequest string = "INS-0"
	RequestTimeout    string = "INS-9"
)

func TopUpWithMpesa(contact string, amount float64, transactionType string) bool {
	res := customerToBusiness(transactionType, contact, amount)
	return res.ResponseCode == SuccessfulRequest
}

func WithdrawWithMpesa(contact string, amount float64, transactionType string) bool {
	res := businessToCustomer(transactionType, contact, amount)
	fmt.Print(res)
	return res.ResponseCode == SuccessfulRequest
}

func customerToBusiness(transactionType string, contact string, amount float64) MpesaResponse {
	if !isValidMpesaContact(contact) {
		log.Panic("Invalid Mpesa contact was received!")
	}

	thirdPartyReference := generateThirdPartyReference(5)
	serviceProviderCode := getMpesaServiceProviderCode()
	url := getMpesaC2BSandboxUrl()

	body := map[string]interface{}{
		"input_TransactionReference": transactionType,
		"input_CustomerMSISDN":       fmt.Sprintf("258%v", contact),
		"input_Amount":               amount,
		"input_ThirdPartyReference":  thirdPartyReference,
		"input_ServiceProviderCode":  serviceProviderCode,
	}

	return initializeRequest(url, body)
}

func businessToCustomer(transactionType string, contact string, amount float64) MpesaResponse {
	if !isValidMpesaContact(contact) {
		log.Panic("Invalid Mpesa contact was received!")
	}

	thirdPartyReference := generateThirdPartyReference(5)
	serviceProviderCode := getMpesaServiceProviderCode()
	url := getMpesaB2CSandboxUrl()

	body := map[string]interface{}{
		"input_TransactionReference": transactionType,
		"input_CustomerMSISDN":       fmt.Sprintf("258%v", contact),
		"input_Amount":               amount,
		"input_ThirdPartyReference":  thirdPartyReference,
		"input_ServiceProviderCode":  serviceProviderCode,
	}

	return initializeRequest(url, body)
}
