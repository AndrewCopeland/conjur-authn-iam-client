package conjur

import (
	"fmt"

	"github.com/cyberark/conjur-api-go/conjurapi"
)

// RetrieveSecret from conjur, if secretID is "" then no error will be returned but value will be (nil. nil)
func RetrieveSecret(config Config, accessToken string, secretID string) ([]byte, error) {
	if secretID == "" {
		return nil, nil
	}

	conjurConfig := conjurapi.Config{
		Account:      config.Account,
		ApplianceURL: config.ApplianceURL,
	}

	client, err := conjurapi.NewClientFromToken(conjurConfig, string(accessToken))
	if err != nil {
		return nil, fmt.Errorf("Failed to create the Conjur client. %s", err)
	}

	value, err := client.RetrieveSecret(secretID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve the variable with ID '%s'. %s", config.SecretID, err)
	}
	return value, nil
}
