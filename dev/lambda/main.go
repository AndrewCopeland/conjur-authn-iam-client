package main

import (
	"context"
	"fmt"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/conjur"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/cyberark/conjur-api-go/conjurapi"
)

func HandleRequest(ctx context.Context) (string, error) {
	config, err := conjur.GetConfig()
	if err != nil {
		return "", err
	}

	accessToken, err := conjur.GetConjurAccessToken(config)
	if err != nil {
		return "", err
	}

	conjurConfig := conjurapi.Config{
		Account:      config.Account,
		ApplianceURL: config.ApplianceURL,
	}

	// Since we have the config and the accessToken lets created out conjurapi.Client
	client, err := conjurapi.NewClientFromToken(conjurConfig, string(accessToken))
	if err != nil {
		return "", fmt.Errorf("Failed to create the Conjur client. %s", err)
	}

	resources, err := client.Resources(nil)
	if err != nil {
		return "", fmt.Errorf("Failed to list resources. %s", err)
	}

	result := ""
	for _, r := range resources {
		id := r["id"].(string)
		result += " - " + id + "\n"
	}

	return result, nil
}

func main() {
	lambda.Start(HandleRequest)
}
