package utils

import (
	"net/http"

	"github.com/cyberark/conjur-api-go/conjurapi"
	"github.com/cyberark/conjur-api-go/conjurapi/authn"
)

func GetConjurHTTPClient(config conjurapi.Config) (*http.Client, error) {
	conjur, err := conjurapi.NewClientFromKey(config,
		authn.LoginPair{
			Login:  "notRealLogin",
			APIKey: "notRealApiKey",
		},
	)
	if err != nil {
		return nil, err
	}

	return conjur.GetHttpClient(), nil
}
