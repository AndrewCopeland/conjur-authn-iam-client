package conjur

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/utils"
	"github.com/cyberark/conjur-api-go/conjurapi"
)

func getAuthnURL(authnURL string, account string, login string) string {
	identifier := url.QueryEscape(login)
	return fmt.Sprintf("%s/%s/%s/authenticate", authnURL, account, identifier)
}

// Authenticate to conjur using the authnURL and conjurAuthnRequest
func Authenticate(authnURL string, login string, conjurAuthnRequest string, config conjurapi.Config) ([]byte, error) {
	httpClient, err := utils.GetConjurHTTPClient(config)
	if err != nil {
		return nil, fmt.Errorf("Failed to init conjur HTTP client. %s", err)
	}

	bodyReader := ioutil.NopCloser(bytes.NewReader([]byte(conjurAuthnRequest)))
	url := getAuthnURL(authnURL, config.Account, login)

	log.Info(log.CAIC005I, url)
	response, err := httpClient.Post(url, "application/json", bodyReader)
	if err != nil {
		return nil, fmt.Errorf("Failed to establish connection to Conjur at url '%s'. %s", url, err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to authenticate to conjur. Recieved status code '%v'", response.StatusCode)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read Conjur Acess Token %s", err)
	}

	return body, err
}
