package api

import (
	"fmt"
	"github.com/wuhoops/silenda/cli/utils"

	"github.com/go-resty/resty/v2"
)

func GetAllEncryptedSecretVariables() {
	//TODO: Call from swagger
	client := resty.New()
	resp, err := client.R().Get(fmt.Sprintf("%s/env", utils.API_URL))
	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Received At: %v", resp.ReceivedAt())
}
