package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/aws"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/conjur"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils"
	"github.com/cyberark/conjur-api-go/conjurapi"
)

func retrieveSecret(config Config, accessToken string) ([]byte, error) {
	client, err := conjurapi.NewClientFromToken(config.Config, string(accessToken))
	if err != nil {
		return nil, fmt.Errorf("Failed to create the Conjur client. %s", err)
	}

	value, err := client.RetrieveSecret(config.SecretID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve the variable with ID '%s'. %s", config.SecretID, err)
	}
	return value, nil
}

func main() {
	config, err := getConfig()
	if err != nil {
		return
	}

	resource, err := aws.GetAwsResource(config.AWSName)
	if err != nil {
		log.Error(log.CAIC002E, err)
		return
	}

	// Retrieveing the AWS IAM Credential
	log.Info(log.CAIC001I, resource.Name())
	credential, err := resource.GetCredential()
	if err != nil {
		log.Error(log.CAIC003E, resource.Name(), err)
		return
	}
	log.Info(log.CAIC002I)

	// Convert the AWS IAM Credential into a Conjur Authentication request
	conjurAuthnRequest, err := utils.GetAuthenticationRequestNow(credential.AccessKeyID, credential.SecretAccessKey, credential.Token)
	if err != nil {
		log.Error(log.CAIC004E, err)
		return
	}
	log.Info(log.CAIC003I)

	// Use the Authentication request to authenticate to Conjur and get a Conjur access token
	log.Info(log.CAIC004I, config.AuthnURL, config.Login, config.Config.Account)
	accessToken, err := conjur.Authenticate(config.AuthnURL, config.Login, conjurAuthnRequest, config.Config)
	if err != nil {
		log.Error(log.CAIC007E, err)
		return
	}
	log.Info(log.CAIC006I, config.Login)

	// Fetch the Secret ID and fail on error
	if config.SecretID != "" {
		value, err := retrieveSecret(config, string(accessToken))
		if err != nil {
			log.Error(log.CAIC011E, config.SecretID, err)
			return
		}
		os.Stdout.Write(value)
	}

	// Write to the Conjur access token file and fail on error
	if config.AccessTokenPath != "" {
		err = ioutil.WriteFile(config.AccessTokenPath, accessToken, 0400)
		if err != nil {
			log.Error(log.CAIC010E, config.AccessTokenPath, err)
			return
		}
		log.Info(log.CAIC007I, config.AccessTokenPath)
	}

	// If secretID and Access token path not provided print out access token to stdout
	if config.SecretID == "" && config.AccessTokenPath == "" {
		fmt.Println(string(accessToken))
	}
}
