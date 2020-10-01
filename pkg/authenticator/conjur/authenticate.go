package conjur

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/AndrewCopeland/conjur-authn-iam-client/pkg/log"
)

func getAuthnURL(authnURL string, account string, login string) string {
	identifier := url.QueryEscape(login)
	return fmt.Sprintf("%s/%s/%s/authenticate", authnURL, account, identifier)
}

// Authenticate to conjur using the authnURL and conjurAuthnRequest
func Authenticate(authnURL string, account string, login string, conjurAuthnRequest string, ignoreSSLVerify bool) ([]byte, error) {
	if ignoreSSLVerify {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	bodyReader := ioutil.NopCloser(bytes.NewReader([]byte(conjurAuthnRequest)))
	url := getAuthnURL(authnURL, account, login)

	log.Info(log.CAIC005I, url)
	response, err := http.Post(url, "", bodyReader)
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
