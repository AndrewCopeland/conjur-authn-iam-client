package conjur

import (
	"flag"
	"os"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
	"github.com/cyberark/conjur-api-go/conjurapi"
)

const (
	ConjurAwsType          = "CONJUR_AWS_TYPE"
	ConjurAuthnUrl         = "CONJUR_AUTHN_URL"
	ConjurAuthnLogin       = "CONJUR_AUTHN_LOGIN"
	ConjurDontAuthenticate = "CONJUR_DONT_AUTHENTICATE"
	ConjurAccessTokenPath  = "CONJUR_ACCESS_TOKEN_PATH"

	FlagAwsType   = "aws-name"
	FlagLogin     = "login"
	FlagAuthnUrl  = "authn-url"
	FlagTokenPath = "token-path"
	FlagSecretID  = "secret"
	FlagSilence   = "silence"

	DescriptionAwsType   = "AWS Resource type name. Environment variable equivalent '" + ConjurAwsType + "'. e.g. ec2, lambda, ecs"
	DescriptionLogin     = "Conjur login that will be used. Environment variable equivalent '" + ConjurAuthnLogin + "'. e.g. host/6634674884744/iam-role-name"
	DescriptionAuthnUrl  = "URL Conjur will be authenticating to. Environment variable equivalent '" + ConjurAuthnUrl + "'. e.g. https://conjur.com/authn-iam/global"
	DescriptionTokenPath = "Write the access token to this file. Environment variable equivalent '" + ConjurAccessTokenPath + "'. e.g. /path/to/access-token.json"
	DescriptionSecretID  = "Retrieve a specific secret from Conjur. e.g. db/postgres/username"
	DescriptionSilence   = "Silence debug and info messages"
)

type Config struct {
	AWSName  string
	Login    string
	AuthnURL string
	Config   conjurapi.Config

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
	login := flag.String(FlagLogin, os.Getenv(ConjurAuthnLogin), DescriptionLogin)
	authnURL := flag.String(FlagAuthnUrl, os.Getenv(ConjurAuthnUrl), DescriptionAuthnUrl)
	config, err := conjurapi.LoadConfig()

	// optional properties
	tokenPath := flag.String(FlagTokenPath, os.Getenv(ConjurAccessTokenPath), DescriptionTokenPath)
	secretID := flag.String(FlagSecretID, "", DescriptionSecretID)
	silence := flag.Bool(FlagSilence, false, DescriptionSilence)
	flag.Parse()

	// Validate mandatory config properties
	if *awsName == "" {
		return Config{}, log.RecordedError(log.CAIC001E, ConjurAwsType, FlagAwsType)
	}

	if *login == "" {
		return Config{}, log.RecordedError(log.CAIC009E, ConjurAuthnLogin, FlagLogin)
	}

	if *authnURL == "" {
		return Config{}, log.RecordedError(log.CAIC006E, ConjurAuthnUrl, FlagAuthnUrl)
	}

	if err != nil {
		return Config{}, log.RecordedError(log.CAIC008E, err)
	}

	if *silence {
		log.EnableSilence()
	}

	return Config{
		AWSName:         *awsName,
		Login:           *login,
		AuthnURL:        *authnURL,
		Config:          config,
		AccessTokenPath: *tokenPath,
		SecretID:        *secretID,
		Silence:         *silence,
	}, nil
}
