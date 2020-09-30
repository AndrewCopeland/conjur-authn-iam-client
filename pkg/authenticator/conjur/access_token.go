package conjur

import (
	"io/ioutil"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/authenticator/aws"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils"
)

// GetConjurAccessToken Get Conjur access token from Conjur
func GetConjurAccessToken(config Config) ([]byte, error) {
	resource, err := aws.GetAwsResource(config.AWSName)
	if err != nil {
		return nil, log.RecordedError(log.CAIC002E, err)
	}

	// Retrieveing the AWS IAM Credential
	log.Info(log.CAIC001I, resource.Name())
	credential, err := resource.GetCredential()
	if err != nil {
		return nil, log.RecordedError(log.CAIC003E, resource.Name(), err)
	}
	log.Info(log.CAIC002I)

	// Convert the AWS IAM Credential into a Conjur Authentication request
	conjurAuthnRequest, err := utils.GetAuthenticationRequestNow(credential.AccessKeyID, credential.SecretAccessKey, credential.Token)
	if err != nil {
		return nil, log.RecordedError(log.CAIC004E, err)
	}
	log.Info(log.CAIC003I)

	// Use the Authentication request to authenticate to Conjur and get a Conjur access token
	log.Info(log.CAIC004I, config.AuthnURL, config.Login, config.Config.Account)
	accessToken, err := Authenticate(config.AuthnURL, config.Login, conjurAuthnRequest, config.Config)
	if err != nil {
		return nil, log.RecordedError(log.CAIC007E, err)
	}
	log.Info(log.CAIC006I, config.Login)

	return accessToken, nil
}

func WriteAccessToken(accessToken []byte, tokenPath string) error {
	if tokenPath == "" {
		return nil
	}

	err := ioutil.WriteFile(tokenPath, accessToken, 0400)
	if err != nil {
		return log.RecordedError(log.CAIC010E, tokenPath, err)
	}
	log.Info(log.CAIC007I, tokenPath)
	return nil
}
