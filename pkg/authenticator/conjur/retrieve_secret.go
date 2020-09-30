package conjur

import (
	"fmt"

	"github.com/cyberark/conjur-api-go/conjurapi"
)

func RetrieveSecret(config Config, accessToken string, secretID string) ([]byte, error) {
	if secretID == "" {
		return nil, nil
	}

	client, err := conjurapi.NewClientFromToken(config.Config, string(accessToken))
	if err != nil {
		return nil, fmt.Errorf("Failed to create the Conjur client. %s", err)
	}

	value, err := client.RetrieveSecret(secretID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve the variable with ID '%s'. %s", config.SecretID, err)
	}
	return value, nil
}
