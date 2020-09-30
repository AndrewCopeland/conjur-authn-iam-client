package main

import (
	"fmt"
	"os"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/conjur"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
)

func main() {
	config, err := conjur.GetConfig()
	if err != nil {
		os.Exit(1)
	}

	accessToken, err := conjur.GetConjurAccessToken(config)
	if err != nil {
		os.Exit(1)
	}

	value, err := conjur.RetrieveSecret(config, string(accessToken), config.SecretID)
	if err != nil {
		log.Error(log.CAIC011E, config.SecretID, err)
		os.Exit(1)
	}
	if value != nil {
		os.Stdout.Write(value)
	}

	err = conjur.WriteAccessToken(accessToken, config.AccessTokenPath)
	if err != nil {
		os.Exit(1)
	}

	// If secretID and Access token path not provided print out access token to stdout
	if config.SecretID == "" && config.AccessTokenPath == "" {
		fmt.Println(string(accessToken))
	}
}
