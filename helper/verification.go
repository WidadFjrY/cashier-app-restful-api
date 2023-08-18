package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
)

func VerifikasiUserEmail(toEmail string, toName string, codeOTP string) {
	url := "https://api.brevo.com/v3/smtp/email"
	apiKey := ""

	// Escape variabel dalam konteks HTML
	escapedToName := template.HTMLEscapeString(toName)
	escapedCodeOTP := template.HTMLEscapeString(codeOTP)

	payload := []byte(`{
		"sender": {
			"name": "Test",
			"email": "widadfjry@gmail.com"
		},
		"to": [
			{
				"email":"` + toEmail + `",
				"name":"` + escapedToName + `"
			}
		],
		"subject": "Test OTP",
		"htmlContent": "<html><head></head><body><p>Hello ` + escapedToName + `,</p>This is OTP for Password reset ` + escapedCodeOTP + `</p></body></html>"
	}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("api-key", apiKey)
	req.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
