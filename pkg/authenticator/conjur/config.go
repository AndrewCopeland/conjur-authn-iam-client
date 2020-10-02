package conjur

import (
	"flag"
	"os"
	"strings"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
)

const (
	ConjurAwsType          = "CONJUR_AWS_TYPE"
	ConjurAccount          = "CONJUR_ACCOUNT"
	ConjurApplianceUrl     = "CONJUR_APPLIANCE_URL"
	ConjurAuthnUrl         = "CONJUR_AUTHN_URL"
	ConjurAuthnLogin       = "CONJUR_AUTHN_LOGIN"
	ConjurDontAuthenticate = "CONJUR_DONT_AUTHENTICATE"
	ConjurAccessTokenPath  = "CONJUR_ACCESS_TOKEN_PATH"
	ConjurIgnoreSSLVerify  = "CONJUR_IGNORE_SSL_VERIFY"

	FlagAwsType         = "aws-name"
	FlagAccount         = "account"
	FlagApplianceUrl    = "url"
	FlagLogin           = "login"
	FlagAuthnUrl        = "authn-url"
	FlagTokenPath       = "token-path"
	FlagSecretID        = "secret"
	FlagSilence         = "silence"
	FlagIgnoreSSLVerify = "ignore-ssl-verify"

	DescriptionAwsType         = "AWS Resource type name. Environment variable equivalent '" + ConjurAwsType + "'. e.g. ec2, lambda, ecs"
	DescriptionAccount         = "The account Conjur has been configued with"
	DescriptionApplianceUrl    = "The URL to the Conjur instance. e.g. https://conjur.com"
	DescriptionLogin           = "Conjur login that will be used. Environment variable equivalent '" + ConjurAuthnLogin + "'. e.g. host/6634674884744/iam-role-name"
	DescriptionAuthnUrl        = "URL Conjur will be authenticating to. Environment variable equivalent '" + ConjurAuthnUrl + "'. e.g. https://conjur.com/authn-iam/global"
	DescriptionTokenPath       = "Write the access token to this file. Environment variable equivalent '" + ConjurAccessTokenPath + "'. e.g. /path/to/access-token.json"
	DescriptionSecretID        = "Retrieve a specific secret from Conjur. e.g. db/postgres/username"
	DescriptionSilence         = "Silence debug and info messages"
	DescriptionIgnoreSSLVerify = "WARNING: Do not verify the SSL certificate provided by Conjur server. THIS SHOULD ONLY BE USED FOR POC"
)

type Config struct {
	AWSName         string
	Account         string
	ApplianceURL    string
	Login           string
	AuthnURL        string
	IgnoreSSLVerify bool

	// If AccessTokenPath & SecretID is not provided then print access token to stdout
	// If only AccessTokenPath is provided then write access token to file
	// If only SecretID is provided then print secret value to stdout
	// If AccessTokenPath & SecretID is provided then write access token to file and print secret value to stdout
	AccessTokenPath string
	SecretID        string
	Silence         bool
}

// Will default to using environment variables if flag is not provided.
// If environment variable and flag is provided then the flag will override the environment variable
func GetConfig() (Config, error) {
	// mandatory properties
	awsName := flag.String(FlagAwsType, os.Getenv(ConjurAwsType), DescriptionAwsType)
	account := flag.String(FlagAccount, os.Getenv(ConjurAccount), DescriptionAccount)
	applianceURL := flag.String(FlagApplianceUrl, os.Getenv(ConjurApplianceUrl), DescriptionApplianceUrl)
	login := flag.String(FlagLogin, os.Getenv(ConjurAuthnLogin), DescriptionLogin)
	authnURL := flag.String(FlagAuthnUrl, os.Getenv(ConjurAuthnUrl), DescriptionAuthnUrl)

	// optional properties
	tokenPath := flag.String(FlagTokenPath, os.Getenv(ConjurAccessTokenPath), DescriptionTokenPath)
	secretID := flag.String(FlagSecretID, "", DescriptionSecretID)
	silence := flag.Bool(FlagSilence, false, DescriptionSilence)

	ignoreStr := strings.ToLower(os.Getenv(ConjurIgnoreSSLVerify))
	ignoreDefault := false
	if ignoreStr == "yes" || ignoreStr == "true" {
		ignoreDefault = true
	}
	ignoreSSLVerify := flag.Bool(FlagIgnoreSSLVerify, ignoreDefault, DescriptionIgnoreSSLVerify)
	flag.Parse()

	// Validate mandatory config properties
	if *awsName == "" {
		return Config{}, log.RecordedError(log.CAIC001E, ConjurAwsType, FlagAwsType)
	}

	if *account == "" {
		return Config{}, log.RecordedError(log.CAIC011E, ConjurAccount, FlagAccount)
	}

	if *applianceURL == "" {
		return Config{}, log.RecordedError(log.CAIC012E, ConjurApplianceUrl, FlagApplianceUrl)
	}

	if *login == "" {
		return Config{}, log.RecordedError(log.CAIC009E, ConjurAuthnLogin, FlagLogin)
	}

	if *authnURL == "" {
		return Config{}, log.RecordedError(log.CAIC006E, ConjurAuthnUrl, FlagAuthnUrl)
	}

	if *silence {
		log.EnableSilence()
	}

	return Config{
		AWSName:         *awsName,
		Account:         *account,
		ApplianceURL:    *applianceURL,
		Login:           *login,
		AuthnURL:        *authnURL,
		AccessTokenPath: *tokenPath,
		SecretID:        *secretID,
		Silence:         *silence,
		IgnoreSSLVerify: *ignoreSSLVerify,
	}, nil
}
