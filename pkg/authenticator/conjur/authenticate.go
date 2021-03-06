package conjur

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/url"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
)

func getAuthnURL(authnURL string, account string, login string) string {
	identifier := url.QueryEscape(login)
	return fmt.Sprintf("%s/%s/%s/authenticate", authnURL, account, identifier)
}

// Authenticate to conjur using the authnURL and conjurAuthnRequest
func Authenticate(authnURL string, account string, login string, conjurAuthnRequest string, ignoreSSLVerify bool, cert []byte) ([]byte, error) {
	client, err := newHTTPSClient(ignoreSSLVerify, cert)
	if err != nil {
		return nil, fmt.Errorf("Failed to create a new HTTPs client. %s", err)
	}

	bodyReader := ioutil.NopCloser(bytes.NewReader([]byte(conjurAuthnRequest)))
	url := getAuthnURL(authnURL, account, login)

	log.Info(log.CAIC005I, url)
	response, err := client.Post(url, "", bodyReader)
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
