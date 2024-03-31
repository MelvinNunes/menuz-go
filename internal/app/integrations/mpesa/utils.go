package mpesa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

var (
	invalidCodeLength = "Invalid code length!"
)

type MpesaResponse struct {
	ResponseCode   string `json:"output_ResponseCode"`
	ResponseDesc   string `json:"output_ResponseDesc"`
	TransactionID  string `json:"output_TransactionID"`
	ConversationID string `json:"output_ConversationID"`
}

func initializeRequest(url string, body map[string]interface{}) MpesaResponse {
	token := getMpesaKey()

	jsonBody, err := json.Marshal(body)

	if err != nil {
		log.Panic("Couldnt parse request body: ", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Panic("Failed to send request to MPESA server!")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", token))
	req.Header.Add("Origin", "developer.mpesa.vm.co.mz")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		log.Print("An error occourred in mpesa response: ", err)
		return MpesaResponse{
			ResponseCode: RequestTimeout,
		}
	}
	defer res.Body.Close()

	var response MpesaResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		log.Panic("Failed to decode response body: ", err)
	}

	return response
}

func generateThirdPartyReference(length int) string {
	if length > 15 || length < 5 {
		return invalidCodeLength
	}

	var code string

	chars := "12346789ABCDEFGHIJKLMNOPQRRSTUVWXYZ987"
	totalCharsIndex := len(chars) - 1

	for i := 0; i < length; i++ {
		position := rand.Intn(totalCharsIndex - 0)
		code = code + string(chars[position])
	}

	return code
}

func isValidMpesaContact(contact string) bool {
	if len(contact) != 9 {
		return false
	}
	if !strings.HasPrefix(contact, "84") && !strings.HasPrefix(contact, "85") {
		return false
	}
	return true
}

func getMpesaKey() string {
	key := os.Getenv("MPESA_KEY")
	if key == "" {
		log.Panic("Please provide a valid MPESA_KEY in your enviroment!")
	}
	return key
}

func getMpesaC2BSandboxUrl() string {
	url := os.Getenv("MPESA_C2B_SANDBOX_URL")
	if url == "" {
		log.Panic("Please provide a valid MPESA_C2B_SANDBOX_URL in your enviroment!")
	}
	return url
}

func getMpesaB2CSandboxUrl() string {
	url := os.Getenv("MPESA_B2C_SANDBOX_URL")
	if url == "" {
		log.Panic("Please provide a valid MPESA_B2C_SANDBOX_URL in your enviroment!")
	}
	return url
}

func getMpesaServiceProviderCode() string {
	url := os.Getenv("MPESA_SERVICE_PROVIDER_CODE")
	if url == "" {
		log.Panic("Please provide a valid MPESA_SERVICE_PROVIDER_CODE in your enviroment!")
	}
	return url
}
