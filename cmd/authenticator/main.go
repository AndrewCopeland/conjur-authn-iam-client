package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/conjur"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
)

func run(config conjur.Config) error {
	accessToken, err := conjur.GetConjurAccessToken(config)
	if err != nil {
		return err
	}

	value, err := conjur.RetrieveSecret(config, string(accessToken), config.SecretID)
	if err != nil {
		return log.RecordedError(log.CAIC011E, config.SecretID, err)
	}
	if value != nil {
		os.Stdout.Write(value)
	}

	err = conjur.WriteAccessToken(accessToken, config.AccessTokenPath)
	if err != nil {
		return err
	}

	// If secretID and Access token path not provided print out access token to stdout
	if config.SecretID == "" && config.AccessTokenPath == "" {
		fmt.Println(string(accessToken))
	}
	return nil
}

func runRetry(config conjur.Config) error {
	for {
		success := false
		var err error
		for i := 0; i < config.Retry; i++ {
			log.Info(log.CAIC008I, i)
			err = run(config)
			if err == nil {
				success = true
				break
			}
			log.Info(log.CAIC009I, config.RetryWait)
			time.Sleep(time.Duration(config.RetryWait) * time.Second)
		}
		if !success {
			return err
		}
		time.Sleep(6 * time.Minute)
	}
}

func main() {
	config, err := conjur.GetConfig()
	if err != nil {
		os.Exit(1)
	}

	// If not refresh then just run once
	if !config.Refresh {
		err := run(config)
		if err != nil {
			os.Exit(1)
		}
		return
	}

	// will run forever if running is successfull.
	// Will return error if failure to authenticate occurs after config.Retry
	err = runRetry(config)
	if err != nil {
		os.Exit(1)
	}
}
